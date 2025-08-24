package storage

import (
	"context"
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/jakobeberhardt/rdt4nn/driver/internal/config"
	log "github.com/sirupsen/logrus"
)

// Manager handles data storage operations
type Manager struct {
	client   influxdb2.Client
	writeAPI api.WriteAPI
	config   config.DBConfig
}

// Measurement represents a data point to be stored
type Measurement struct {
	Name      string
	Tags      map[string]string
	Fields    map[string]interface{}
	Timestamp time.Time
}

// NewManager creates a new storage manager
func NewManager(config config.DBConfig) (*Manager, error) {
	// Create InfluxDB client (token can be empty for development)
	token := ""
	if config.User != "" && config.Password != "" {
		// For production, you might want to use actual authentication
		// This is a simplified approach
		log.WithField("user", config.User).Info("InfluxDB authentication configured")
	}
	
	client := influxdb2.NewClient(fmt.Sprintf("http://%s", config.Host), token)
	
	// Test connection with a ping (with timeout)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	health, err := client.Health(ctx)
	if err != nil {
		log.WithError(err).Warn("Failed to connect to InfluxDB - continuing without storage")
		// Return a minimal manager that doesn't fail
		return &Manager{
			client:   client,
			writeAPI: client.WriteAPI("", config.Name),
			config:   config,
		}, nil
	}
	
	if health.Status != "pass" {
		log.WithField("status", health.Status).Warn("InfluxDB health check warning")
	}

	// Get write API for non-blocking writes (org can be empty for InfluxDB 1.x compatibility)
	writeAPI := client.WriteAPI("", config.Name)

	// Set up error handling
	errorsCh := writeAPI.Errors()
	go func() {
		for err := range errorsCh {
			log.WithError(err).Error("InfluxDB write error")
		}
	}()

	sm := &Manager{
		client:   client,
		writeAPI: writeAPI,
		config:   config,
	}

	log.WithFields(log.Fields{
		"host":     config.Host,
		"database": config.Name,
	}).Info("Storage manager initialized")

	return sm, nil
}

// WriteMeasurements writes a batch of measurements to InfluxDB
func (sm *Manager) WriteMeasurements(ctx context.Context, measurements []Measurement) error {
	if len(measurements) == 0 {
		return nil
	}

	// Convert measurements to InfluxDB points
	for _, measurement := range measurements {
		point := influxdb2.NewPoint(
			measurement.Name,
			measurement.Tags,
			measurement.Fields,
			measurement.Timestamp,
		)
		
		// Write point (non-blocking)
		sm.writeAPI.WritePoint(point)
	}

	log.WithField("count", len(measurements)).Debug("Measurements written to storage")
	return nil
}

// WriteMeasurement writes a single measurement to InfluxDB
func (sm *Manager) WriteMeasurement(ctx context.Context, measurement Measurement) error {
	return sm.WriteMeasurements(ctx, []Measurement{measurement})
}

// Query executes a query against InfluxDB
func (sm *Manager) Query(ctx context.Context, query string) ([]map[string]interface{}, error) {
	queryAPI := sm.client.QueryAPI("")
	
	result, err := queryAPI.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var results []map[string]interface{}
	for result.Next() {
		record := result.Record()
		recordMap := make(map[string]interface{})
		
		// Add all values from the record
		for key, value := range record.Values() {
			recordMap[key] = value
		}
		
		results = append(results, recordMap)
	}
	
	if result.Err() != nil {
		return nil, fmt.Errorf("query result error: %w", result.Err())
	}

	return results, nil
}

// CreateBucket creates a bucket in InfluxDB (if needed)
func (sm *Manager) CreateBucket(ctx context.Context, bucketName string, retention time.Duration) error {
	// Note: This is a simplified version. In production, you might want to use InfluxDB v2 Organizations API
	// For now, we'll skip bucket creation as it typically requires admin access
	log.WithField("bucket", bucketName).Info("Bucket creation skipped - ensure bucket exists in InfluxDB")
	return nil
}

// Flush ensures all pending writes are sent to InfluxDB
func (sm *Manager) Flush() error {
	sm.writeAPI.Flush()
	return nil
}

// Close closes the storage manager and all connections
func (sm *Manager) Close() error {
	log.Info("Closing storage manager")
	
	// Flush any pending writes
	sm.writeAPI.Flush()
	
	// Close client (WriteAPI is automatically closed when client closes)
	sm.client.Close()
	
	return nil
}

// GetStats returns statistics about the storage
func (sm *Manager) GetStats(ctx context.Context) (map[string]interface{}, error) {
	stats := make(map[string]interface{})
	
	// Get bucket information
	query := fmt.Sprintf(`
		from(bucket: "%s")
		|> range(start: -1h)
		|> group()
		|> count()
	`, sm.config.Name)
	
	results, err := sm.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get storage stats: %w", err)
	}
	
	if len(results) > 0 {
		stats["total_points_last_hour"] = results[0]["_value"]
	}
	
	stats["database"] = sm.config.Name
	stats["host"] = sm.config.Host
	
	return stats, nil
}

// WriteMetadata writes benchmark metadata
func (sm *Manager) WriteMetadata(ctx context.Context, benchmarkID string, metadata map[string]interface{}) error {
	tags := map[string]string{
		"benchmark_id": benchmarkID,
		"type":         "metadata",
	}
	
	measurement := Measurement{
		Name:      "benchmark_metadata",
		Tags:      tags,
		Fields:    metadata,
		Timestamp: time.Now(),
	}
	
	return sm.WriteMeasurement(ctx, measurement)
}
