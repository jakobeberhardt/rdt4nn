package benchmark

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/jakobeberhardt/rdt4nn/driver/internal/config"
	"github.com/jakobeberhardt/rdt4nn/driver/internal/container"
	"github.com/jakobeberhardt/rdt4nn/driver/internal/profiler"
	"github.com/jakobeberhardt/rdt4nn/driver/internal/scheduler"
	"github.com/jakobeberhardt/rdt4nn/driver/internal/storage"
	log "github.com/sirupsen/logrus"
)

// Runner orchestrates the entire benchmark execution
type Runner struct {
	config         *config.Config
	containerMgr   *container.Manager
	profilerMgr    *profiler.Manager
	schedulerMgr   scheduler.Scheduler
	storageMgr     *storage.Manager
	benchmarkID    string
	startTime      time.Time
	ctx            context.Context
	cancel         context.CancelFunc
	wg             sync.WaitGroup
}

// NewRunner creates a new benchmark runner with the given configuration
func NewRunner(cfg *config.Config) (*Runner, error) {
	ctx, cancel := context.WithCancel(context.Background())
	
	benchmarkID := fmt.Sprintf("%s_%d", cfg.Benchmark.Name, time.Now().Unix())
	
	// Initialize container manager
	containerMgr, err := container.NewManager(cfg.Benchmark.Docker.Auth)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("failed to create container manager: %w", err)
	}

	// Initialize storage manager
	storageMgr, err := storage.NewManager(cfg.Benchmark.Data.DB)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("failed to create storage manager: %w", err)
	}

	// Initialize profiler manager
	profilerMgr, err := profiler.NewManager(&cfg.Benchmark.Data, benchmarkID, storageMgr)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("failed to create profiler manager: %w", err)
	}

	// Initialize scheduler
	schedulerMgr, err := scheduler.NewScheduler(cfg.Benchmark.Scheduler.Implementation, cfg.Benchmark.Scheduler.RDT)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("failed to create scheduler: %w", err)
	}

	return &Runner{
		config:       cfg,
		containerMgr: containerMgr,
		profilerMgr:  profilerMgr,
		schedulerMgr: schedulerMgr,
		storageMgr:   storageMgr,
		benchmarkID:  benchmarkID,
		ctx:          ctx,
		cancel:       cancel,
	}, nil
}

// Run executes the complete benchmark lifecycle
func (r *Runner) Run() error {
	r.startTime = time.Now()
	
	log.WithFields(log.Fields{
		"benchmark_id": r.benchmarkID,
		"start_time":   r.startTime,
	}).Info("Starting benchmark execution")

	// Set up signal handling for graceful shutdown
	r.setupSignalHandling()

	// Phase 1: Preparation
	if err := r.prepare(); err != nil {
		return fmt.Errorf("preparation phase failed: %w", err)
	}

	// Phase 2: Execution
	if err := r.execute(); err != nil {
		return fmt.Errorf("execution phase failed: %w", err)
	}

	// Phase 3: Cleanup
	if err := r.cleanup(); err != nil {
		log.WithError(err).Error("Cleanup phase encountered errors")
		return err
	}

	return nil
}

// prepare handles the preparation phase
func (r *Runner) prepare() error {
	log.Info("Entering preparation phase")

	// Pull all required images
	if err := r.containerMgr.PullImages(r.ctx, r.config.Container); err != nil {
		return fmt.Errorf("failed to pull images: %w", err)
	}

	// Create all containers
	if err := r.containerMgr.CreateContainers(r.ctx, r.config.Container, r.benchmarkID); err != nil {
		return fmt.Errorf("failed to create containers: %w", err)
	}

	// Initialize scheduler
	if err := r.schedulerMgr.Initialize(r.ctx); err != nil {
		return fmt.Errorf("failed to initialize scheduler: %w", err)
	}

	// Initialize profiler
	if err := r.profilerMgr.Initialize(r.ctx); err != nil {
		return fmt.Errorf("failed to initialize profiler: %w", err)
	}

	log.Info("Preparation phase completed successfully")
	return nil
}

// execute handles the main execution phase
func (r *Runner) execute() error {
	log.Info("Entering execution phase")

	// Start profiling
	r.wg.Add(1)
	go func() {
		defer r.wg.Done()
		if err := r.profilerMgr.StartProfiling(r.ctx); err != nil {
			log.WithError(err).Error("Profiling failed")
		}
	}()

	// Start scheduler monitoring
	r.wg.Add(1)
	go func() {
		defer r.wg.Done()
		if err := r.schedulerMgr.Monitor(r.ctx, r.containerMgr.GetContainerIDs()); err != nil {
			log.WithError(err).Error("Scheduler monitoring failed")
		}
	}()

	// Start containers according to schedule
	if err := r.startContainersScheduled(); err != nil {
		return fmt.Errorf("failed to start containers: %w", err)
	}

	// Wait for benchmark completion or manual termination
	r.waitForCompletion()

	log.Info("Execution phase completed")
	return nil
}

// startContainersScheduled starts containers according to their scheduled start times
func (r *Runner) startContainersScheduled() error {
	// Create a map of start times to containers
	startSchedule := make(map[int][]string)
	for name, containerCfg := range r.config.Container {
		startSchedule[containerCfg.Start] = append(startSchedule[containerCfg.Start], name)
	}

	// Start containers in scheduled order
	for startTime := 0; ; startTime++ {
		select {
		case <-r.ctx.Done():
			return r.ctx.Err()
		default:
		}

		if containers, exists := startSchedule[startTime]; exists {
			log.WithFields(log.Fields{
				"time":       startTime,
				"containers": containers,
			}).Info("Starting scheduled containers")

			for _, containerName := range containers {
				if err := r.containerMgr.StartContainer(r.ctx, containerName); err != nil {
					return fmt.Errorf("failed to start container %s: %w", containerName, err)
				}

				// Apply scheduler policies
				containerID := r.containerMgr.GetContainerID(containerName)
				if containerID != "" {
					if err := r.schedulerMgr.ScheduleContainer(r.ctx, containerID, r.config.Container[containerName]); err != nil {
						log.WithError(err).WithField("container", containerName).Warn("Failed to apply scheduler policy")
					}
				}
			}
		}

		// Check if we need to continue
		if len(startSchedule) == 0 || startTime > r.getMaxStartTime() {
			break
		}

		// Wait for next second
		select {
		case <-r.ctx.Done():
			return r.ctx.Err()
		case <-time.After(time.Second):
		}
	}

	return nil
}

// waitForCompletion waits for benchmark completion based on configuration
func (r *Runner) waitForCompletion() {
	if r.config.Benchmark.MaxT == -1 {
		log.Info("Benchmark set to run indefinitely, waiting for manual termination")
		<-r.ctx.Done()
	} else {
		duration := time.Duration(r.config.Benchmark.MaxT) * time.Second
		log.WithField("duration", duration).Info("Waiting for benchmark completion")
		
		select {
		case <-r.ctx.Done():
			log.Info("Benchmark terminated manually")
		case <-time.After(duration):
			log.Info("Benchmark completed after specified duration")
			r.cancel()
		}
	}

	// Wait for all goroutines to finish
	r.wg.Wait()
}

// cleanup handles the cleanup phase
func (r *Runner) cleanup() error {
	log.Info("Entering cleanup phase")

	var errors []error

	// Stop profiling
	if err := r.profilerMgr.Stop(); err != nil {
		errors = append(errors, fmt.Errorf("failed to stop profiler: %w", err))
	}

	// Stop and remove containers
	if err := r.containerMgr.StopAndCleanup(r.ctx); err != nil {
		errors = append(errors, fmt.Errorf("failed to cleanup containers: %w", err))
	}

	// Finalize scheduler
	if err := r.schedulerMgr.Finalize(r.ctx); err != nil {
		errors = append(errors, fmt.Errorf("failed to finalize scheduler: %w", err))
	}

	// Close storage connections
	if err := r.storageMgr.Close(); err != nil {
		errors = append(errors, fmt.Errorf("failed to close storage: %w", err))
	}

	if len(errors) > 0 {
		for _, err := range errors {
			log.WithError(err).Error("Cleanup error")
		}
		return fmt.Errorf("cleanup completed with %d errors", len(errors))
	}

	log.Info("Cleanup phase completed successfully")
	return nil
}

// setupSignalHandling sets up graceful shutdown on interrupt signals
func (r *Runner) setupSignalHandling() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigChan
		log.WithField("signal", sig).Info("Received termination signal, initiating graceful shutdown")
		r.cancel()
	}()
}

// getMaxStartTime returns the maximum start time from all containers
func (r *Runner) getMaxStartTime() int {
	maxStartTime := 0
	for _, containerCfg := range r.config.Container {
		if containerCfg.Start > maxStartTime {
			maxStartTime = containerCfg.Start
		}
	}
	return maxStartTime
}
