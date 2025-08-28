package profiler

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/jakobeberhardt/rdt4nn/driver/internal/storage"
	log "github.com/sirupsen/logrus"
)

// PerfCollector collects hardware performance counters using perf with cgroup filtering
type PerfCollector struct {
	benchmarkID  string
	containerIDs map[string]string // container name -> container ID mapping
	samplingRate int               // sampling rate in milliseconds
}

// NewPerfCollector creates a new perf collector
func NewPerfCollector(benchmarkID string) *PerfCollector {
	return &PerfCollector{
		benchmarkID:  benchmarkID,
		containerIDs: make(map[string]string),
		samplingRate: 100, // default 100ms, will be updated from config
	}
}

// SetContainerIDs sets the container IDs for monitoring
func (p *PerfCollector) SetContainerIDs(containerIDs map[string]string) {
	p.containerIDs = containerIDs
}

// SetSamplingRate sets the sampling rate from configuration
func (p *PerfCollector) SetSamplingRate(rateMs int) {
	p.samplingRate = rateMs
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

// getContainerPID gets the PID of a container
func (p *PerfCollector) getContainerPID(containerID string) (int, error) {
	cmd := exec.Command("docker", "inspect", "-f", "{{.State.Pid}}", containerID)
	output, err := cmd.Output()
	if err != nil {
		return 0, fmt.Errorf("failed to get container PID: %w", err)
	}

	pidStr := strings.TrimSpace(string(output))
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		return 0, fmt.Errorf("invalid PID: %w", err)
	}

	return pid, nil
}

// getContainerCgroup gets the cgroup path for a container
func (p *PerfCollector) getContainerCgroup(pid int) (string, error) {
	cgroupFile := fmt.Sprintf("/proc/%d/cgroup", pid)
	file, err := os.Open(cgroupFile)
	if err != nil {
		return "", fmt.Errorf("failed to open cgroup file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Look for the unified cgroup (cgroup v2) or the cpu,cpuacct cgroup
		if strings.HasPrefix(line, "0:") {
			parts := strings.Split(line, ":")
			if len(parts) >= 3 {
				return parts[2], nil
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("failed to read cgroup file: %w", err)
	}

	return "", fmt.Errorf("cgroup not found for PID %d", pid)
}

// Collect collects performance counters using perf with cgroup filtering
func (p *PerfCollector) Collect(ctx context.Context, timestamp time.Time) ([]storage.Measurement, error) {
	var measurements []storage.Measurement

	// Define hardware performance events to collect based on your POC
	events := []string{
		"cycles",           // cpu-cycles
		"instructions",     // instructions
		"cache-misses",     // cache-misses
		"cache-references", // cache-references
		"branches",         // branch-instructions
		"branch-misses",    // branch-misses
		"bus-cycles",       // bus-cycles
		"ref-cycles",       // ref-cycles
		"stalled-cycles-frontend", // stalled-cycles-frontend
		// Cache events
		"L1-dcache-loads",
		"L1-dcache-load-misses",
		"L1-dcache-stores", 
		"L1-dcache-store-misses",
		"L1-icache-load-misses",
		"LLC-loads",
		"LLC-load-misses",
		"LLC-stores",
		"LLC-store-misses",
	}

	// Collect metrics for each container
	for containerName, containerID := range p.containerIDs {
		if containerID == "" {
			continue
		}

		// Get container PID
		pid, err := p.getContainerPID(containerID)
		if err != nil {
			log.WithError(err).WithField("container", containerName).Debug("Failed to get container PID, skipping perf collection")
			continue
		}

		// Get container cgroup
		cgroup, err := p.getContainerCgroup(pid)
		if err != nil {
			log.WithError(err).WithField("container", containerName).Debug("Failed to get container cgroup, skipping perf collection")
			continue
		}

		// Run perf stat with cgroup filtering  
		// Create a timeout context specifically for perf to avoid cancellation issues
		perfCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		
		cmd := exec.CommandContext(perfCtx, "sudo", "perf", "stat", 
			"-e", strings.Join(events, ","),
			"-a", 
			"--cgroup", cgroup,
			"--", "sleep", "0.01") // 10ms measurement duration

		log.WithFields(log.Fields{
			"container": containerName,
			"cgroup":    cgroup,
			"events":    len(events),
		}).Debug("Running perf stat command")

		output, err := cmd.CombinedOutput()
		if err != nil {
			log.WithError(err).WithFields(log.Fields{
				"container": containerName,
				"cgroup":    cgroup,
				"output":    string(output),
			}).Debug("Failed to run perf stat, skipping perf collection for container")
			continue
		}

		log.WithFields(log.Fields{
			"container":   containerName,
			"output_size": len(output),
		}).Debug("Perf stat completed successfully")

		// Parse the output and create measurements
		containerMeasurements := p.parsePerfOutput(string(output), timestamp, containerName, containerID)
		measurements = append(measurements, containerMeasurements...)
	}

	return measurements, nil
}

// parsePerfOutput parses perf stat output and converts to measurements
// The output format from perf stat is: count unit event [additional info]
func (p *PerfCollector) parsePerfOutput(output string, timestamp time.Time, containerName, containerID string) []storage.Measurement {
	var measurements []storage.Measurement
	
	lines := strings.Split(output, "\n")
	perfData := make(map[string]uint64)
	
	sampleLines := lines
	if len(lines) > 5 {
		sampleLines = lines[:5]
	}
	log.WithFields(log.Fields{
		"container":    containerName,
		"total_lines":  len(lines),
		"sample_lines": sampleLines,
	}).Debug("Parsing perf output")
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, " Performance counter stats") ||
		   strings.HasPrefix(line, "Performance counter stats") || strings.Contains(line, "seconds time elapsed") ||
		   strings.Contains(line, "Some events weren't counted") || strings.Contains(line, "echo") ||
		   strings.Contains(line, "perf stat") {
			continue
		}

		// Parse standard perf stat output format
		// Examples:
		// 62,314,106      cycles                    /system.slice/docker-...
		// 122,567,399     instructions              /system.slice/docker-...
		// <not supported> cache-misses:u
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}

		countStr := strings.TrimSpace(parts[0])
		if strings.HasPrefix(countStr, "<not") || countStr == "" {
			// Set unsupported counters to 0 as requested
			continue
		}

		// Remove comma separators and parse count
		countStr = strings.ReplaceAll(countStr, ",", "")
		count, err := strconv.ParseUint(countStr, 10, 64)
		if err != nil {
			log.WithError(err).WithField("line", line).Debug("Failed to parse perf counter value")
			continue
		}

		// Extract event name (remove :u suffix if present)
		eventName := strings.TrimSpace(parts[1])
		eventName = strings.Split(eventName, ":")[0] // Remove :u, :k suffixes
		
		// Map event names to our standard names
		switch eventName {
		case "cycles", "cpu-cycles":
			perfData["cpu_cycles"] = count
		case "instructions":
			perfData["instructions"] = count
		case "cache-references":
			perfData["cache_references"] = count
		case "cache-misses":
			perfData["cache_misses"] = count
		case "branches", "branch-instructions":
			perfData["branch_instructions"] = count
		case "branch-misses":
			perfData["branch_misses"] = count
		case "bus-cycles":
			perfData["bus_cycles"] = count
		case "ref-cycles":
			perfData["ref_cycles"] = count
		case "stalled-cycles-frontend", "idle-cycles-frontend":
			perfData["stalled_cycles_frontend"] = count
		case "L1-dcache-loads":
			perfData["l1_dcache_loads"] = count
		case "L1-dcache-load-misses":
			perfData["l1_dcache_load_misses"] = count
		case "L1-dcache-stores":
			perfData["l1_dcache_stores"] = count
		case "L1-dcache-store-misses":
			perfData["l1_dcache_store_misses"] = count
		case "L1-icache-load-misses":
			perfData["l1_icache_load_misses"] = count
		case "LLC-loads":
			perfData["llc_loads"] = count
		case "LLC-load-misses":
			perfData["llc_load_misses"] = count
		case "LLC-stores":
			perfData["llc_stores"] = count
		case "LLC-store-misses":
			perfData["llc_store_misses"] = count
		}
	}

	log.WithFields(log.Fields{
		"container":      containerName,
		"metrics_parsed": len(perfData),
		"metrics":        perfData,
	}).Debug("Perf data parsed")

	// Create measurements for each counter
	for metricName, value := range perfData {
		tags := map[string]string{
			"benchmark_id":  p.benchmarkID,
			"collector":     "perf",
			"container":     containerName,
			"container_id":  containerID,
			"metric":        metricName,
		}

		measurements = append(measurements, storage.Measurement{
			Name:      "perf_counter",
			Tags:      tags,
			Fields:    map[string]interface{}{"value": value},
			Timestamp: timestamp,
		})
	}

	// Create additional measurements for calculated metrics (if we have the base metrics)
	if cycles, hasCycles := perfData["cpu_cycles"]; hasCycles {
		if instructions, hasInstr := perfData["instructions"]; hasInstr && cycles > 0 {
			ipc := float64(instructions) / float64(cycles)
			tags := map[string]string{
				"benchmark_id":  p.benchmarkID,
				"collector":     "perf",
				"container":     containerName,
				"container_id":  containerID,
				"metric":        "ipc",
			}
			measurements = append(measurements, storage.Measurement{
				Name:      "perf_counter",
				Tags:      tags,
				Fields:    map[string]interface{}{"value": ipc},
				Timestamp: timestamp,
			})
		}
	}

	if cacheRefs, hasRefs := perfData["cache_references"]; hasRefs {
		if cacheMisses, hasMisses := perfData["cache_misses"]; hasMisses && cacheRefs > 0 {
			missRate := float64(cacheMisses) / float64(cacheRefs)
			tags := map[string]string{
				"benchmark_id":  p.benchmarkID,
				"collector":     "perf",
				"container":     containerName,
				"container_id":  containerID,
				"metric":        "cache_miss_rate",
			}
			measurements = append(measurements, storage.Measurement{
				Name:      "perf_counter",
				Tags:      tags,
				Fields:    map[string]interface{}{"value": missRate},
				Timestamp: timestamp,
			})
		}
	}

	return measurements
}

func (p *PerfCollector) Close() error {
	return nil
}

func (p *PerfCollector) Name() string {
	return "perf"
}
