package profiler

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/jakobeberhardt/rdt4nn/driver/internal/storage"
	log "github.com/sirupsen/logrus"
)

type DockerStatsCollector struct {
	benchmarkID  string
	client       *client.Client
	containerIDs map[string]string // name -> ID mapping
}

func NewDockerStatsCollector(benchmarkID string) *DockerStatsCollector {
	return &DockerStatsCollector{
		benchmarkID:  benchmarkID,
		containerIDs: make(map[string]string),
	}
}

func (d *DockerStatsCollector) SetContainerIDs(containerIDs map[string]string) {
	d.containerIDs = make(map[string]string)
	for name, id := range containerIDs {
		d.containerIDs[name] = id
	}
}

func (d *DockerStatsCollector) Initialize(ctx context.Context) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("failed to create Docker client: %w", err)
	}
	d.client = cli
	return nil
}

func (d *DockerStatsCollector) Collect(ctx context.Context, timestamp time.Time) ([]storage.Measurement, error) {
	filterArgs := filters.NewArgs()
	filterArgs.Add("label", "benchmark.id="+d.benchmarkID)
	
	containers, err := d.client.ContainerList(ctx, types.ContainerListOptions{
		Filters: filterArgs,
		All:     false, // Only running containers
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list containers: %w", err)
	}

	var measurements []storage.Measurement

	for _, container := range containers {
		// Get container stats
		stats, err := d.client.ContainerStats(ctx, container.ID, false)
		if err != nil {
			log.WithError(err).WithField("container", container.ID[:12]).Warn("Failed to get container stats")
			continue
		}

		var statsJSON types.StatsJSON
		if err := json.NewDecoder(stats.Body).Decode(&statsJSON); err != nil {
			stats.Body.Close()
			log.WithError(err).WithField("container", container.ID[:12]).Warn("Failed to decode container stats")
			continue
		}
		stats.Body.Close()

		containerMeasurements := d.convertStatsToMeasurements(container, statsJSON, timestamp)
		measurements = append(measurements, containerMeasurements...)
	}

	return measurements, nil
}

func (d *DockerStatsCollector) convertStatsToMeasurements(container types.Container, stats types.StatsJSON, timestamp time.Time) []storage.Measurement {
	containerName := ""
	containerIndex := ""
	
	if name, exists := container.Labels["benchmark.name"]; exists {
		containerName = name
	}
	if index, exists := container.Labels["benchmark.index"]; exists {
		containerIndex = index
	}

	tags := map[string]string{
		"benchmark_id":     d.benchmarkID,
		"container_id":     container.ID[:12],
		"container_name":   containerName,
		"container_index":  containerIndex,
		"collector":        "docker_stats",
	}

	var measurements []storage.Measurement

	cpuUsage := float64(0)
	if stats.PreCPUStats.CPUUsage.TotalUsage > 0 {
		cpuDelta := float64(stats.CPUStats.CPUUsage.TotalUsage - stats.PreCPUStats.CPUUsage.TotalUsage)
		systemDelta := float64(stats.CPUStats.SystemUsage - stats.PreCPUStats.SystemUsage)
		if systemDelta > 0 {
			cpuUsage = (cpuDelta / systemDelta) * float64(len(stats.CPUStats.CPUUsage.PercpuUsage)) * 100.0
		}
	}

	measurements = append(measurements, storage.Measurement{
		Name:      "cpu_usage_percent",
		Tags:      copyTags(tags),
		Fields:    map[string]interface{}{"value": cpuUsage},
		Timestamp: timestamp,
	})

	memUsage := float64(stats.MemoryStats.Usage)
	memLimit := float64(stats.MemoryStats.Limit)
	memPercent := float64(0)
	if memLimit > 0 {
		memPercent = (memUsage / memLimit) * 100.0
	}

	measurements = append(measurements, storage.Measurement{
		Name:      "memory_usage_bytes",
		Tags:      copyTags(tags),
		Fields:    map[string]interface{}{"value": memUsage},
		Timestamp: timestamp,
	})

	measurements = append(measurements, storage.Measurement{
		Name:      "memory_usage_percent",
		Tags:      copyTags(tags),
		Fields:    map[string]interface{}{"value": memPercent},
		Timestamp: timestamp,
	})

	if len(stats.Networks) > 0 {
		var rxBytes, txBytes uint64
		for _, network := range stats.Networks {
			rxBytes += network.RxBytes
			txBytes += network.TxBytes
		}

		measurements = append(measurements, storage.Measurement{
			Name:      "network_rx_bytes",
			Tags:      copyTags(tags),
			Fields:    map[string]interface{}{"value": float64(rxBytes)},
			Timestamp: timestamp,
		})

		measurements = append(measurements, storage.Measurement{
			Name:      "network_tx_bytes",
			Tags:      copyTags(tags),
			Fields:    map[string]interface{}{"value": float64(txBytes)},
			Timestamp: timestamp,
		})
	}

	if len(stats.BlkioStats.IoServiceBytesRecursive) > 0 {
		var readBytes, writeBytes uint64
		for _, blkio := range stats.BlkioStats.IoServiceBytesRecursive {
			if blkio.Op == "Read" {
				readBytes += blkio.Value
			} else if blkio.Op == "Write" {
				writeBytes += blkio.Value
			}
		}

		measurements = append(measurements, storage.Measurement{
			Name:      "blkio_read_bytes",
			Tags:      copyTags(tags),
			Fields:    map[string]interface{}{"value": float64(readBytes)},
			Timestamp: timestamp,
		})

		measurements = append(measurements, storage.Measurement{
			Name:      "blkio_write_bytes",
			Tags:      copyTags(tags),
			Fields:    map[string]interface{}{"value": float64(writeBytes)},
			Timestamp: timestamp,
		})
	}

	return measurements
}

func copyTags(original map[string]string) map[string]string {
	copy := make(map[string]string)
	for k, v := range original {
		copy[k] = v
	}
	return copy
}

func (d *DockerStatsCollector) Close() error {
	if d.client != nil {
		return d.client.Close()
	}
	return nil
}

func (d *DockerStatsCollector) Name() string {
	return "docker_stats"
}
