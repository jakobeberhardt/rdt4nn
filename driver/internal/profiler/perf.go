package profiler

import (
	"context"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/jakobeberhardt/rdt4nn/driver/internal/storage"
	log "github.com/sirupsen/logrus"
)

// PerfCollector collects hardware performance counters using perf
type PerfCollector struct {
	benchmarkID string
	perfCmd     *exec.Cmd
	perfPID     int
}

// NewPerfCollector creates a new perf collector
func NewPerfCollector(benchmarkID string) *PerfCollector {
	return &PerfCollector{
		benchmarkID: benchmarkID,
	}
}

// Initialize initializes the perf collector
func (p *PerfCollector) Initialize(ctx context.Context) error {
	// Check if perf is available
	if _, err := exec.LookPath("perf"); err != nil {
		return fmt.Errorf("perf command not found: %w", err)
	}

	log.Info("Perf collector initialized")
	return nil
}

// Collect collects performance counters using perf
func (p *PerfCollector) Collect(ctx context.Context, timestamp time.Time) ([]storage.Measurement, error) {
	// Define performance events to collect
	events := []string{
		"cycles",
		"instructions",
		"cache-references",
		"cache-misses",
		"branch-instructions",
		"branch-misses",
		"page-faults",
		"context-switches",
		"cpu-migrations",
		"L1-dcache-loads",
		"L1-dcache-load-misses",
		"LLC-loads",
		"LLC-load-misses",
	}

	// Run perf stat for system-wide collection
	cmd := exec.CommandContext(ctx, "perf", "stat", "-x", ",", "-e", strings.Join(events, ","), "--", "sleep", "0.001")
	
	output, err := cmd.CombinedOutput()
	if err != nil {
		// perf might not be available or might need privileges
		log.WithError(err).Debug("Failed to run perf stat, skipping perf collection")
		return []storage.Measurement{}, nil
	}

	return p.parsePerfOutput(string(output), timestamp), nil
}

// parsePerfOutput parses perf stat output and converts to measurements
func (p *PerfCollector) parsePerfOutput(output string, timestamp time.Time) []storage.Measurement {
	var measurements []storage.Measurement
	
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		// Parse CSV format: value,unit,event,time,percent
		parts := strings.Split(line, ",")
		if len(parts) < 3 {
			continue
		}

		valueStr := strings.TrimSpace(parts[0])
		if valueStr == "<not supported>" || valueStr == "<not counted>" {
			continue
		}

		// Remove any formatting from the value
		valueStr = strings.ReplaceAll(valueStr, ",", "")
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			continue
		}

		event := strings.TrimSpace(parts[2])
		
		tags := map[string]string{
			"benchmark_id": p.benchmarkID,
			"collector":    "perf",
			"event":        event,
		}

		measurements = append(measurements, storage.Measurement{
			Name:      "perf_counter",
			Tags:      tags,
			Fields:    map[string]interface{}{"value": value},
			Timestamp: timestamp,
		})
	}

	return measurements
}

// Close closes the perf collector
func (p *PerfCollector) Close() error {
	// Nothing specific to close for perf
	return nil
}

// Name returns the collector name
func (p *PerfCollector) Name() string {
	return "perf"
}
