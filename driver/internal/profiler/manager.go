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
	config           *config.DataConfig
	benchmarkID      string
	benchmarkIDNum   int64
	storage          *storage.Manager
	collectors       []Collector
	ticker           *time.Ticker
	stopChan         chan struct{}
	wg               sync.WaitGroup
	containerConfigs map[string]*config.ContainerConfig // Add container configs
	schedulerType    string                              // Add scheduler type
}

// ProfilerManager defines the interface that all profiler managers must implement
type ProfilerManager interface {
	Initialize(ctx context.Context, containerIDs map[string]string) error
	StartProfiling(ctx context.Context, containerIDs map[string]string) error
	Stop() error
}

// Collector defines the interface for different profiling collectors
type Collector interface {
	Initialize(ctx context.Context) error
	Collect(ctx context.Context, timestamp time.Time) ([]storage.Measurement, error)
	Close() error
	Name() string
}

// NewManager creates a new profiler manager
func NewManager(config *config.DataConfig, benchmarkID string, benchmarkIDNum int64, storage *storage.Manager) (*Manager, error) {
	pm := &Manager{
		config:         config,
		benchmarkID:    benchmarkID,
		benchmarkIDNum: benchmarkIDNum,
		storage:        storage,
		collectors:     make([]Collector, 0),
		stopChan:       make(chan struct{}),
	}

	// Initialize collectors based on configuration
	var dockerCollector *DockerStatsCollector
	var perfCollector *PerfCollector
	var rdtCollector *RDTCollector
	
	if config.DockerStats {
		dockerCollector = NewDockerStatsCollector(benchmarkID)
		pm.collectors = append(pm.collectors, dockerCollector)
	}

	if config.Perf {
		perfCollector = NewPerfCollector(benchmarkID)
		pm.collectors = append(pm.collectors, perfCollector)
	}

	if config.RDT {
		rdtCollector = NewRDTCollector(benchmarkID)
		pm.collectors = append(pm.collectors, rdtCollector)
	}

	// Add a comprehensive collector that aggregates data from all sources
	// This ensures benchmark metrics are written with proper data
	if config.DockerStats || config.Perf || config.RDT {
		comprehensiveCollector := NewDataAggregatorCollector(
			benchmarkID, 
			benchmarkIDNum, 
			storage, 
			dockerCollector, 
			perfCollector, 
			rdtCollector,
		)
		pm.collectors = append(pm.collectors, comprehensiveCollector)
	}

	if len(pm.collectors) == 0 {
		return nil, fmt.Errorf("no profiling collectors configured")
	}

	log.WithField("collectors", len(pm.collectors)).Info("Profiler manager initialized")
	return pm, nil
}

// Initialize initializes all collectors
func (pm *Manager) Initialize(ctx context.Context, containerIDs map[string]string) error {
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
func (pm *Manager) StartProfiling(ctx context.Context, containerIDs map[string]string) error {
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

// DataAggregatorCollector collects data from all sources and writes comprehensive benchmark metrics
type DataAggregatorCollector struct {
	benchmarkID     string
	benchmarkIDNum  int64
	storage         *storage.Manager
	startTime       time.Time
	samplingStep    int64
	dockerCollector *DockerStatsCollector
	perfCollector   *PerfCollector
	rdtCollector    *RDTCollector
}

// NewDataAggregatorCollector creates a new data aggregator collector
func NewDataAggregatorCollector(
	benchmarkID string, 
	benchmarkIDNum int64, 
	storage *storage.Manager,
	dockerCollector *DockerStatsCollector,
	perfCollector *PerfCollector,
	rdtCollector *RDTCollector,
) *DataAggregatorCollector {
	return &DataAggregatorCollector{
		benchmarkID:     benchmarkID,
		benchmarkIDNum:  benchmarkIDNum,
		storage:         storage,
		dockerCollector: dockerCollector,
		perfCollector:   perfCollector,
		rdtCollector:    rdtCollector,
	}
}

// Initialize initializes the data aggregator collector
func (dac *DataAggregatorCollector) Initialize(ctx context.Context) error {
	dac.startTime = time.Now()
	dac.samplingStep = 0
	return nil
}

// Collect aggregates data from all sources and writes comprehensive benchmark metrics
func (dac *DataAggregatorCollector) Collect(ctx context.Context, timestamp time.Time) ([]storage.Measurement, error) {
	dac.samplingStep++
	
	// Collect data from all available sources
	var dockerMeasurements []storage.Measurement
	var perfMeasurements []storage.Measurement
	var rdtMeasurements []storage.Measurement
	
	if dac.dockerCollector != nil {
		dockerMeasurements, _ = dac.dockerCollector.Collect(ctx, timestamp)
	}
	
	if dac.perfCollector != nil {
		perfMeasurements, _ = dac.perfCollector.Collect(ctx, timestamp)
	}
	
	if dac.rdtCollector != nil {
		rdtMeasurements, _ = dac.rdtCollector.Collect(ctx, timestamp)
	}
	
	// Aggregate data by container and write comprehensive metrics
	containerData := dac.aggregateDataByContainer(dockerMeasurements, perfMeasurements, rdtMeasurements)
	
	// Write comprehensive metrics for each container
	for containerInfo, data := range containerData {
		relativeTime := timestamp.Sub(dac.startTime).Milliseconds()
		
		metrics := &storage.BenchmarkMetrics{
			// Benchmark metadata
			BenchmarkID:       dac.benchmarkIDNum,
			BenchmarkStarted:  dac.startTime,
			SamplingFrequency: 500, // TODO: Get this from config
			UsedScheduler:     "default", // TODO: Get this from scheduler
			
			// Container metadata
			ContainerName:  containerInfo,
			ContainerIndex: 0, // TODO: Get from container config
			ContainerImage: "unknown", // TODO: Get from container config  
			ContainerCore:  0, // TODO: Get from container config
			
			// Timing information
			UTCTimestamp:  timestamp,
			RelativeTime:  relativeTime,
			SamplingStep:  dac.samplingStep,
			CPUExecutedOn: dac.getCurrentCPU(),
			
			// Performance data
			DockerMetrics: data.DockerData,
			PerfMetrics:   data.PerfData,
			RDTMetrics:    data.RDTData,
		}
		
		// Write comprehensive metrics to storage
		if err := dac.storage.WriteBenchmarkMetrics(ctx, metrics); err != nil {
			log.WithError(err).WithField("container", containerInfo).Error("Failed to write comprehensive metrics")
		}
	}
	
	// Return empty measurements since the writing is handled above
	return []storage.Measurement{}, nil
}

// ContainerData holds aggregated data for a container
type ContainerData struct {
	DockerData *storage.DockerData
	PerfData   *storage.PerfData
	RDTData    *storage.RDTData
}

// aggregateDataByContainer organizes measurements by container
func (dac *DataAggregatorCollector) aggregateDataByContainer(
	dockerMeasurements, perfMeasurements, rdtMeasurements []storage.Measurement,
) map[string]*ContainerData {
	containerData := make(map[string]*ContainerData)
	
	// Process Docker measurements
	for _, measurement := range dockerMeasurements {
		containerName := dac.extractContainerName(measurement)
		if containerName == "" {
			continue
		}
		
		if _, exists := containerData[containerName]; !exists {
			containerData[containerName] = &ContainerData{}
		}
		
		// Convert measurement to DockerData
		dockerData := dac.convertToDockerData(measurement)
		containerData[containerName].DockerData = dockerData
	}
	
	// Process Perf measurements
	for _, measurement := range perfMeasurements {
		containerName := dac.extractContainerName(measurement)
		if containerName == "" {
			continue
		}
		
		if _, exists := containerData[containerName]; !exists {
			containerData[containerName] = &ContainerData{}
		}
		
		// Convert measurement to PerfData
		perfData := dac.convertToPerfData(measurement)
		containerData[containerName].PerfData = perfData
	}
	
	// Process RDT measurements
	for _, measurement := range rdtMeasurements {
		containerName := dac.extractContainerName(measurement)
		if containerName == "" {
			continue
		}
		
		if _, exists := containerData[containerName]; !exists {
			containerData[containerName] = &ContainerData{}
		}
		
		// Convert measurement to RDTData
		rdtData := dac.convertToRDTData(measurement)
		containerData[containerName].RDTData = rdtData
	}
	
	return containerData
}

// Helper methods to extract and convert data
func (dac *DataAggregatorCollector) extractContainerName(measurement storage.Measurement) string {
	if containerName, exists := measurement.Tags["container"]; exists {
		return containerName
	}
	if containerName, exists := measurement.Tags["container_name"]; exists {
		return containerName
	}
	return ""
}

func (dac *DataAggregatorCollector) convertToDockerData(measurement storage.Measurement) *storage.DockerData {
	dockerData := &storage.DockerData{}
	
	// Extract CPU metrics
	if val, ok := measurement.Fields["cpu_usage_percent"]; ok {
		if f, ok := val.(float64); ok {
			dockerData.CPUUsagePercent = f
		}
	}
	
	// Extract memory metrics
	if val, ok := measurement.Fields["memory_usage"]; ok {
		if u, ok := val.(uint64); ok {
			dockerData.MemoryUsage = u
		}
	}
	
	// Add more field extractions as needed
	return dockerData
}

func (dac *DataAggregatorCollector) convertToPerfData(measurement storage.Measurement) *storage.PerfData {
	perfData := &storage.PerfData{}
	
	// Extract perf metrics
	if val, ok := measurement.Fields["cpu_cycles"]; ok {
		if u, ok := val.(uint64); ok {
			perfData.CPUCycles = u
		}
	}
	
	if val, ok := measurement.Fields["instructions"]; ok {
		if u, ok := val.(uint64); ok {
			perfData.Instructions = u
		}
	}
	
	// Add more field extractions as needed
	return perfData
}

func (dac *DataAggregatorCollector) convertToRDTData(measurement storage.Measurement) *storage.RDTData {
	rdtData := &storage.RDTData{}
	
	// Extract RDT metrics
	if val, ok := measurement.Fields["llc_occupancy"]; ok {
		if u, ok := val.(uint64); ok {
			rdtData.LLCOccupancy = u
		}
	}
	
	// Add more field extractions as needed
	return rdtData
}

func (dac *DataAggregatorCollector) getCurrentCPU() int {
	// TODO: Implement actual CPU detection
	return 0
}

// Close closes the data aggregator collector
func (dac *DataAggregatorCollector) Close() error {
	// No cleanup needed
	return nil
}

// Name returns the name of this collector
func (dac *DataAggregatorCollector) Name() string {
	return "data_aggregator"
}
