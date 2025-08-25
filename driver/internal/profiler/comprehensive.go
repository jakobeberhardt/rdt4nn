package profiler

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/jakobeberhardt/rdt4nn/driver/internal/config"
	"github.com/jakobeberhardt/rdt4nn/driver/internal/storage"
	log "github.com/sirupsen/logrus"
)

// ComprehensiveManager coordinates all profiling and creates comprehensive metrics
type ComprehensiveManager struct {
	config         *config.DataConfig
	benchmarkID    string
	benchmarkIDNum int64
	startTime      time.Time
	endTime        time.Time
	storage        *storage.Manager
	collectors     []Collector
	ticker         *time.Ticker
	stopChan       chan struct{}
	wg             sync.WaitGroup
	samplingStep   int64 
	
	containerConfigs map[string]*config.ContainerConfig
	schedulerType    string
}

func NewComprehensiveManager(
	config *config.DataConfig,
	benchmarkID string,
	benchmarkIDNum int64,
	startTime time.Time,
	storage *storage.Manager,
	containerConfigs map[string]*config.ContainerConfig,
	schedulerType string,
) (*ComprehensiveManager, error) {
	pm := &ComprehensiveManager{
		config:           config,
		benchmarkID:      benchmarkID,
		benchmarkIDNum:   benchmarkIDNum,
		startTime:        startTime,
		storage:          storage,
		collectors:       make([]Collector, 0),
		stopChan:         make(chan struct{}),
		samplingStep:     0,
		containerConfigs: containerConfigs,
		schedulerType:    schedulerType,
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

	log.WithField("collectors", len(pm.collectors)).Info("Comprehensive profiler manager initialized")
	return pm, nil
}

func (pm *ComprehensiveManager) Initialize(ctx context.Context, containerIDs map[string]string) error {
	log.Info("Initializing profiler collectors")

	for _, collector := range pm.collectors {
		if err := collector.Initialize(ctx); err != nil {
			return fmt.Errorf("failed to initialize collector %s: %w", collector.Name(), err)
		}
		
		if dockerCollector, ok := collector.(*DockerStatsCollector); ok {
			dockerCollector.SetContainerIDs(containerIDs)
		}
		
		log.WithField("collector", collector.Name()).Info("Collector initialized")
	}

	return nil
}

func (pm *ComprehensiveManager) StartProfiling(ctx context.Context, containerIDs map[string]string) error {
	log.WithField("frequency_ms", pm.config.ProfileFrequency).Info("Starting comprehensive profiling")

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
			go pm.collectComprehensiveMetrics(ctx, timestamp, containerIDs)
		}
	}
}

func (pm *ComprehensiveManager) collectComprehensiveMetrics(ctx context.Context, timestamp time.Time, containerIDs map[string]string) {
	defer pm.wg.Done()

	pm.samplingStep++
	step := pm.samplingStep

	relativeTime := timestamp.Sub(pm.startTime).Milliseconds()
	
	currentCPU := runtime.NumCPU() 

	dockerData := make(map[string]*storage.DockerData)
	perfData := make(map[string]*storage.PerfData)
	rdtData := make(map[string]*storage.RDTData)

	for _, collector := range pm.collectors {
		measurements, err := collector.Collect(ctx, timestamp)
		if err != nil {
			log.WithError(err).WithField("collector", collector.Name()).Debug("Failed to collect metrics")
			continue
		}

		switch collector.Name() {
		case "docker_stats":
			dockerData = pm.parseDockerMeasurements(measurements)
		case "perf":
			perfData = pm.parsePerfMeasurements(measurements)
		case "rdt":
			rdtData = pm.parseRDTMeasurements(measurements)
		}
	}

	for containerName, containerID := range containerIDs {
		containerConfig, exists := pm.containerConfigs[containerName]
		if !exists {
			log.WithField("container", containerName).Warn("Container config not found")
			continue
		}

		shortContainerID := containerID[:12]
		dockerMetrics := dockerData[shortContainerID] 
		
		log.WithFields(log.Fields{
			"container_name":       containerName,
			"full_container_id":    containerID,
			"short_container_id":   shortContainerID,
			"docker_data_found":    dockerMetrics != nil,
			"available_docker_keys": fmt.Sprintf("%v", getKeys(dockerData)),
		}).Debug("Looking up Docker metrics for container")

		metrics := &storage.BenchmarkMetrics{
			BenchmarkID:       pm.benchmarkIDNum,
			BenchmarkStarted:  pm.startTime,
			BenchmarkFinished: pm.endTime, 
			SamplingFrequency: pm.config.ProfileFrequency,
			UsedScheduler:     pm.schedulerType,

			ContainerName:  containerName,
			ContainerIndex: containerConfig.Index,
			ContainerImage: containerConfig.Image,
			ContainerCore:  containerConfig.Core,

			UTCTimestamp:  timestamp,
			RelativeTime:  relativeTime,
			SamplingStep:  step,
			CPUExecutedOn: currentCPU,

			DockerMetrics: dockerMetrics, // Use the corrected lookup
			PerfMetrics:   perfData[containerName], // Perf might use container name
			RDTMetrics:    rdtData[containerName],  // RDT might use container name
		}

		if err := pm.storage.WriteBenchmarkMetrics(ctx, metrics); err != nil {
			log.WithError(err).WithField("container", containerName).Error("Failed to write comprehensive metrics")
		}
	}
}

func (pm *ComprehensiveManager) parseDockerMeasurements(measurements []storage.Measurement) map[string]*storage.DockerData {
	result := make(map[string]*storage.DockerData)
	
	// Group measurements by container ID
	for _, measurement := range measurements {
		containerID := measurement.Tags["container_id"]
		if containerID == "" {
			continue
		}

		// Initialize docker data for this container if not exists
		if _, exists := result[containerID]; !exists {
			result[containerID] = &storage.DockerData{}
		}
		
		dockerData := result[containerID]
		
		// Parse based on measurement name
		switch measurement.Name {
		case "cpu_usage_percent":
			if val, ok := measurement.Fields["value"].(float64); ok {
				dockerData.CPUUsagePercent = val
			}
		case "cpu_usage_total":
			if val, ok := measurement.Fields["value"].(float64); ok {
				dockerData.CPUUsageTotal = uint64(val)
			}
		case "cpu_usage_kernel":
			if val, ok := measurement.Fields["value"].(float64); ok {
				dockerData.CPUUsageKernel = uint64(val)
			}
		case "cpu_usage_user":
			if val, ok := measurement.Fields["value"].(float64); ok {
				dockerData.CPUUsageUser = uint64(val)
			}
		case "cpu_throttling":
			if val, ok := measurement.Fields["value"].(float64); ok {
				dockerData.CPUThrottling = uint64(val)
			}
		case "memory_usage_bytes":
			if val, ok := measurement.Fields["value"].(float64); ok {
				dockerData.MemoryUsage = uint64(val)
			}
		case "memory_limit":
			if val, ok := measurement.Fields["value"].(float64); ok {
				dockerData.MemoryLimit = uint64(val)
			}
		case "memory_usage_percent":
			if val, ok := measurement.Fields["value"].(float64); ok {
				dockerData.MemoryUsagePercent = val
			}
		case "memory_cache":
			if val, ok := measurement.Fields["value"].(float64); ok {
				dockerData.MemoryCache = uint64(val)
			}
		case "memory_rss":
			if val, ok := measurement.Fields["value"].(float64); ok {
				dockerData.MemoryRSS = uint64(val)
			}
		case "memory_swap":
			if val, ok := measurement.Fields["value"].(float64); ok {
				dockerData.MemorySwap = uint64(val)
			}
		case "network_rx_bytes":
			if val, ok := measurement.Fields["value"].(float64); ok {
				dockerData.NetworkRxBytes = uint64(val)
			}
		case "network_tx_bytes":
			if val, ok := measurement.Fields["value"].(float64); ok {
				dockerData.NetworkTxBytes = uint64(val)
			}
		case "network_rx_packets":
			if val, ok := measurement.Fields["value"].(float64); ok {
				dockerData.NetworkRxPackets = uint64(val)
			}
		case "network_tx_packets":
			if val, ok := measurement.Fields["value"].(float64); ok {
				dockerData.NetworkTxPackets = uint64(val)
			}
		case "disk_read_bytes":
			if val, ok := measurement.Fields["value"].(float64); ok {
				dockerData.DiskReadBytes = uint64(val)
			}
		case "disk_write_bytes":
			if val, ok := measurement.Fields["value"].(float64); ok {
				dockerData.DiskWriteBytes = uint64(val)
			}
		case "disk_read_ops":
			if val, ok := measurement.Fields["value"].(float64); ok {
				dockerData.DiskReadOps = uint64(val)
			}
		case "disk_write_ops":
			if val, ok := measurement.Fields["value"].(float64); ok {
				dockerData.DiskWriteOps = uint64(val)
			}
		}
	}
	
	return result
}

// Helper function to get keys from a map for debugging
func getKeys(m map[string]*storage.DockerData) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func (pm *ComprehensiveManager) parsePerfMeasurements(measurements []storage.Measurement) map[string]*storage.PerfData {
	result := make(map[string]*storage.PerfData)
	
	for _, measurement := range measurements {
		containerName := measurement.Tags["container_name"]
		if containerName == "" {
			continue
		}

		perfData := &storage.PerfData{}
		
		if val, ok := measurement.Fields["cpu_cycles"].(uint64); ok {
			perfData.CPUCycles = val
		}
		if val, ok := measurement.Fields["instructions"].(uint64); ok {
			perfData.Instructions = val
		}
		if val, ok := measurement.Fields["cache_references"].(uint64); ok {
			perfData.CacheReferences = val
		}
		if val, ok := measurement.Fields["cache_misses"].(uint64); ok {
			perfData.CacheMisses = val
		}
		// Calculate derived metrics
		if perfData.Instructions > 0 && perfData.CPUCycles > 0 {
			perfData.IPC = float64(perfData.Instructions) / float64(perfData.CPUCycles)
		}
		if perfData.CacheReferences > 0 && perfData.CacheMisses > 0 {
			perfData.CacheMissRate = float64(perfData.CacheMisses) / float64(perfData.CacheReferences)
		}

		result[containerName] = perfData
	}
	
	return result
}

func (pm *ComprehensiveManager) parseRDTMeasurements(measurements []storage.Measurement) map[string]*storage.RDTData {
	result := make(map[string]*storage.RDTData)
	
	for _, measurement := range measurements {
		containerName := measurement.Tags["container_name"]
		if containerName == "" {
			continue
		}

		rdtData := &storage.RDTData{}
		
		// Parse RDT fields from measurement
		if val, ok := measurement.Fields["llc_occupancy"].(uint64); ok {
			rdtData.LLCOccupancy = val
		}
		if val, ok := measurement.Fields["local_mem_bw"].(float64); ok {
			rdtData.LocalMemBW = val
		}
		if val, ok := measurement.Fields["remote_mem_bw"].(float64); ok {
			rdtData.RemoteMemBW = val
		}
		if val, ok := measurement.Fields["total_mem_bw"].(float64); ok {
			rdtData.TotalMemBW = val
		}

		result[containerName] = rdtData
	}
	
	return result
}

func (pm *ComprehensiveManager) SetEndTime(endTime time.Time) {
	pm.endTime = endTime
}

// Stop stops the profiling
func (pm *ComprehensiveManager) Stop() error {
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
		return fmt.Errorf("encountered %d errors while stopping collectors", len(errors))
	}
	return nil
}
