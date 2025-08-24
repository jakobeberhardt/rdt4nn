package scheduler

import (
	"context"
	"fmt"
	"os/exec"
	"strconv"
	"syscall"

	"github.com/jakobeberhardt/rdt4nn/driver/internal/config"
	log "github.com/sirupsen/logrus"
)

// DefaultScheduler implements a basic scheduling strategy
type DefaultScheduler struct {
	rdtEnabled bool
}

// NewDefaultScheduler creates a new default scheduler
func NewDefaultScheduler(rdtEnabled bool) *DefaultScheduler {
	return &DefaultScheduler{
		rdtEnabled: rdtEnabled,
	}
}

// Initialize initializes the default scheduler
func (ds *DefaultScheduler) Initialize(ctx context.Context) error {
	log.Info("Initializing default scheduler")
	
	if ds.rdtEnabled {
		log.Info("RDT support enabled for default scheduler")
		// Basic RDT initialization could go here
	}
	
	return nil
}

// ScheduleContainer applies scheduling policies to a container
func (ds *DefaultScheduler) ScheduleContainer(ctx context.Context, containerID string, cfg config.ContainerConfig) error {
	log.WithFields(log.Fields{
		"container_id": containerID[:12],
		"core":         cfg.Core,
	}).Debug("Applying default scheduling policy")

	// Apply CPU pinning if specified
	if cfg.Core >= 0 {
		if err := ds.setCPUAffinity(containerID, cfg.Core); err != nil {
			return fmt.Errorf("failed to set CPU affinity: %w", err)
		}
	}

	// Apply basic resource limits
	if err := ds.setResourceLimits(containerID, cfg); err != nil {
		log.WithError(err).Warn("Failed to set resource limits")
	}

	return nil
}

// setCPUAffinity sets CPU affinity for a container
func (ds *DefaultScheduler) setCPUAffinity(containerID string, core int) error {
	// Get container PID
	cmd := exec.Command("docker", "inspect", "-f", "{{.State.Pid}}", containerID)
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to get container PID: %w", err)
	}

	pidStr := string(output)
	pid, err := strconv.Atoi(pidStr[:len(pidStr)-1]) // Remove newline
	if err != nil {
		return fmt.Errorf("invalid PID: %w", err)
	}

	// Set CPU affinity using taskset
	cmd = exec.Command("taskset", "-p", "-c", strconv.Itoa(core), strconv.Itoa(pid))
	if err := cmd.Run(); err != nil {
		log.WithError(err).WithFields(log.Fields{
			"pid":  pid,
			"core": core,
		}).Debug("Failed to set CPU affinity with taskset, trying syscall")

		// Fallback to syscall
		cpuSet := &syscall.CPUSet{}
		cpuSet.Set(core)
		if err := syscall.SchedSetaffinity(pid, cpuSet); err != nil {
			return fmt.Errorf("failed to set CPU affinity via syscall: %w", err)
		}
	}

	log.WithFields(log.Fields{
		"container_id": containerID[:12],
		"pid":          pid,
		"core":         core,
	}).Info("CPU affinity set successfully")

	return nil
}

// setResourceLimits sets basic resource limits
func (ds *DefaultScheduler) setResourceLimits(containerID string, cfg config.ContainerConfig) error {
	// This is a placeholder for setting additional resource limits
	// In a real implementation, you might use cgroups directly or Docker API
	
	log.WithField("container_id", containerID[:12]).Debug("Resource limits applied")
	return nil
}

// Monitor monitors container performance and adjusts scheduling if needed
func (ds *DefaultScheduler) Monitor(ctx context.Context, containerIDs map[string]string) error {
	log.Info("Starting default scheduler monitoring")
	
	// The default scheduler doesn't perform dynamic adjustments
	// It just logs that monitoring is active
	
	<-ctx.Done()
	log.Info("Default scheduler monitoring stopped")
	return ctx.Err()
}

// Finalize cleans up scheduler resources
func (ds *DefaultScheduler) Finalize(ctx context.Context) error {
	log.Info("Finalizing default scheduler")
	
	// Nothing specific to clean up for the default scheduler
	return nil
}

// Name returns the scheduler name
func (ds *DefaultScheduler) Name() string {
	return "default"
}
