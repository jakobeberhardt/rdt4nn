package profiler

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/jakobeberhardt/rdt4nn/driver/internal/storage"
	log "github.com/sirupsen/logrus"
)

// PerfData represents buffered perf measurements for a container
type PerfBuffer struct {
	containerName string
	data         map[string]uint64
	timestamp    time.Time
	mutex        sync.RWMutex
}

// PerfCollector collects hardware performance counters using perf with cgroup filtering
type PerfCollector struct {
	benchmarkID  string
	containerIDs map[string]string // container name -> container ID mapping
	samplingRate int               // sampling rate in milliseconds
	
	// Buffering system
	buffers     map[string]*PerfBuffer // container name -> latest perf data
	bufferMutex sync.RWMutex
	stopChan    chan struct{}
	wg          sync.WaitGroup
	running     bool
}

// NewPerfCollector creates a new perf collector
func NewPerfCollector(benchmarkID string) *PerfCollector {
	return &PerfCollector{
		benchmarkID:  benchmarkID,
		containerIDs: make(map[string]string),
		buffers:      make(map[string]*PerfBuffer),
		samplingRate: 100, // default 100ms, will be updated from config
		stopChan:     make(chan struct{}),
	}
}

// SetContainerIDs sets the container IDs for monitoring
func (p *PerfCollector) SetContainerIDs(containerIDs map[string]string) {
	p.containerIDs = containerIDs
	
	// Initialize buffers for each container
	p.bufferMutex.Lock()
	defer p.bufferMutex.Unlock()
	
	for containerName := range containerIDs {
		p.buffers[containerName] = &PerfBuffer{
			containerName: containerName,
			data:         make(map[string]uint64),
		}
	}
}

// SetSamplingRate sets the sampling rate from configuration
func (p *PerfCollector) SetSamplingRate(rateMs int) {
	if rateMs > 0 {
		p.samplingRate = rateMs
	}
}

// Initialize initializes the perf collector and starts the background collection
func (p *PerfCollector) Initialize(ctx context.Context) error {
	// Check if perf is available
	if _, err := exec.LookPath("perf"); err != nil {
		return fmt.Errorf("perf command not found: %w", err)
	}

	// Start background collection goroutine
	p.running = true
	p.wg.Add(1)
	go p.backgroundCollection(ctx)

	log.WithField("sampling_rate_ms", p.samplingRate).Info("Perf collector initialized with background collection")
	return nil
}

// backgroundCollection runs perf collection in a background goroutine at high frequency
func (p *PerfCollector) backgroundCollection(ctx context.Context) {
	defer p.wg.Done()
	
	ticker := time.NewTicker(time.Duration(p.samplingRate) * time.Millisecond)
	defer ticker.Stop()
	
	log.WithField("interval_ms", p.samplingRate).Debug("Starting background perf collection")
	
	for {
		select {
		case <-ctx.Done():
			log.Debug("Background perf collection stopped due to context cancellation")
			return
		case <-p.stopChan:
			log.Debug("Background perf collection stopped")
			return
		case timestamp := <-ticker.C:
			p.collectAndBuffer(timestamp)
		}
	}
}

// collectAndBuffer collects perf data and updates buffers
func (p *PerfCollector) collectAndBuffer(timestamp time.Time) {
	// Define hardware performance events to collect
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
			// Don't log every failure to reduce noise
			continue
		}

		// Get container cgroup
		cgroup, err := p.getContainerCgroup(pid)
		if err != nil {
			continue
		}

		// Run perf stat with cgroup filtering  
		// Create a short timeout for high-frequency collection
		perfCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		
		cmd := exec.CommandContext(perfCtx, "sudo", "perf", "stat", 
			"-e", strings.Join(events, ","),
			"-a", 
			"--cgroup", cgroup,
			"--", "sleep", "0.01") // 10ms measurement duration

		output, err := cmd.CombinedOutput()
		cancel()
		
		if err != nil {
			// Continue to next container on error
			continue
		}

		// Parse and update buffer
		perfData := p.parsePerfOutputToMap(string(output))
		if len(perfData) > 0 {
			p.bufferMutex.Lock()
			if buffer, exists := p.buffers[containerName]; exists {
				buffer.mutex.Lock()
				buffer.data = perfData
				buffer.timestamp = timestamp
				buffer.mutex.Unlock()
			}
			p.bufferMutex.Unlock()
			
			log.WithFields(log.Fields{
				"container": containerName,
				"metrics":   len(perfData),
			}).Debug("Updated perf buffer")
		}
	}
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

// Collect returns the latest buffered performance data
func (p *PerfCollector) Collect(ctx context.Context, timestamp time.Time) ([]storage.Measurement, error) {
	var measurements []storage.Measurement

	p.bufferMutex.RLock()
	defer p.bufferMutex.RUnlock()

	// Get latest buffered data for each container
	for containerName, containerID := range p.containerIDs {
		if containerID == "" {
			continue
		}

		buffer, exists := p.buffers[containerName]
		if !exists {
			continue
		}

		buffer.mutex.RLock()
		perfData := make(map[string]uint64)
		bufferTimestamp := buffer.timestamp
		// Copy data to avoid holding lock too long
		for k, v := range buffer.data {
			perfData[k] = v
		}
		buffer.mutex.RUnlock()

		// Skip if no data available or data is too old (more than 2 seconds)
		if len(perfData) == 0 || time.Since(bufferTimestamp) > 2*time.Second {
			log.WithFields(log.Fields{
				"container": containerName,
				"data_age": time.Since(bufferTimestamp),
				"has_data": len(perfData) > 0,
			}).Debug("Skipping stale or missing perf data")
			continue
		}

		// Convert buffered data to measurements
		containerMeasurements := p.convertToMeasurements(perfData, timestamp, containerName, containerID)
		measurements = append(measurements, containerMeasurements...)

		log.WithFields(log.Fields{
			"container": containerName,
			"metrics":   len(perfData),
			"data_age":  time.Since(bufferTimestamp),
		}).Debug("Using buffered perf data")
	}

	return measurements, nil
}

// convertToMeasurements converts perf data map to measurement format
func (p *PerfCollector) convertToMeasurements(perfData map[string]uint64, timestamp time.Time, containerName, containerID string) []storage.Measurement {
	var measurements []storage.Measurement

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

// parsePerfOutputToMap parses perf stat output and returns a map of metrics
func (p *PerfCollector) parsePerfOutputToMap(output string) map[string]uint64 {
	perfData := make(map[string]uint64)
	lines := strings.Split(output, "\n")
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, " Performance counter stats") ||
		   strings.HasPrefix(line, "Performance counter stats") || strings.Contains(line, "seconds time elapsed") ||
		   strings.Contains(line, "Some events weren't counted") || strings.Contains(line, "echo") ||
		   strings.Contains(line, "perf stat") {
			continue
		}

		// Parse standard perf stat output format
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}

		countStr := strings.TrimSpace(parts[0])
		if strings.HasPrefix(countStr, "<not") || countStr == "" {
			continue
		}

		// Remove comma separators and parse count
		countStr = strings.ReplaceAll(countStr, ",", "")
		count, err := strconv.ParseUint(countStr, 10, 64)
		if err != nil {
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
	
	return perfData
}

// Close closes the perf collector and stops background collection
func (p *PerfCollector) Close() error {
	if p.running {
		p.running = false
		close(p.stopChan)
		p.wg.Wait()
		log.Debug("Perf collector background collection stopped")
	}
	return nil
}

// Name returns the collector name
func (p *PerfCollector) Name() string {
	return "perf"
}
