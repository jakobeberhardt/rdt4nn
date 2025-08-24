package profiler

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/jakobeberhardt/rdt4nn/driver/internal/config"
	"github.com/jakobeberhardt/rdt4nn/driver/internal/storage"
	log "github.com/sirupsen/logrus"
)

// Manager coordinates all profiling activities
type Manager struct {
	config      *config.DataConfig
	benchmarkID string
	storage     *storage.Manager
	collectors  []Collector
	ticker      *time.Ticker
	stopChan    chan struct{}
	wg          sync.WaitGroup
}

// Collector defines the interface for different profiling collectors
type Collector interface {
	Initialize(ctx context.Context) error
	Collect(ctx context.Context, timestamp time.Time) ([]storage.Measurement, error)
	Close() error
	Name() string
}

// NewManager creates a new profiler manager
func NewManager(config *config.DataConfig, benchmarkID string, storage *storage.Manager) (*Manager, error) {
	pm := &Manager{
		config:      config,
		benchmarkID: benchmarkID,
		storage:     storage,
		collectors:  make([]Collector, 0),
		stopChan:    make(chan struct{}),
	}

	// Initialize collectors based on configuration
	if config.DockerStats {
		dockerCollector := NewDockerStatsCollector(benchmarkID)
		pm.collectors = append(pm.collectors, dockerCollector)
	}

	if config.Perf {
		perfCollector := NewPerfCollector(benchmarkID)
		pm.collectors = append(pm.collectors, perfCollector)
	}

	if config.RDT {
		rdtCollector := NewRDTCollector(benchmarkID)
		pm.collectors = append(pm.collectors, rdtCollector)
	}

	if len(pm.collectors) == 0 {
		return nil, fmt.Errorf("no profiling collectors configured")
	}

	log.WithField("collectors", len(pm.collectors)).Info("Profiler manager initialized")
	return pm, nil
}

// Initialize initializes all collectors
func (pm *Manager) Initialize(ctx context.Context) error {
	log.Info("Initializing profiler collectors")

	for _, collector := range pm.collectors {
		if err := collector.Initialize(ctx); err != nil {
			return fmt.Errorf("failed to initialize collector %s: %w", collector.Name(), err)
		}
		log.WithField("collector", collector.Name()).Info("Collector initialized")
	}

	return nil
}

// StartProfiling starts the profiling loop
func (pm *Manager) StartProfiling(ctx context.Context) error {
	log.WithField("frequency_ms", pm.config.ProfileFrequency).Info("Starting profiling")

	duration := time.Duration(pm.config.ProfileFrequency) * time.Millisecond
	pm.ticker = time.NewTicker(duration)
	defer pm.ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Info("Profiling stopped due to context cancellation")
			return ctx.Err()
		case <-pm.stopChan:
			log.Info("Profiling stopped")
			return nil
		case timestamp := <-pm.ticker.C:
			pm.wg.Add(1)
			go pm.collectMetrics(ctx, timestamp)
		}
	}
}

// collectMetrics collects metrics from all collectors
func (pm *Manager) collectMetrics(ctx context.Context, timestamp time.Time) {
	defer pm.wg.Done()

	var allMeasurements []storage.Measurement

	for _, collector := range pm.collectors {
		measurements, err := collector.Collect(ctx, timestamp)
		if err != nil {
			log.WithError(err).WithField("collector", collector.Name()).Error("Failed to collect metrics")
			continue
		}
		allMeasurements = append(allMeasurements, measurements...)
	}

	if len(allMeasurements) > 0 {
		if err := pm.storage.WriteMeasurements(ctx, allMeasurements); err != nil {
			log.WithError(err).Error("Failed to write measurements to storage")
		}
	}
}

// Stop stops the profiling
func (pm *Manager) Stop() error {
	log.Info("Stopping profiler")

	close(pm.stopChan)
	pm.wg.Wait()

	var errors []error
	for _, collector := range pm.collectors {
		if err := collector.Close(); err != nil {
			errors = append(errors, fmt.Errorf("failed to close collector %s: %w", collector.Name(), err))
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("encountered %d errors while stopping profiler", len(errors))
	}

	return nil
}
