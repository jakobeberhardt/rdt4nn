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
	// Create InfluxDB client
	client := influxdb2.NewClient(fmt.Sprintf("http://%s", config.Host), "")
	
	// Test connection with a ping
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	health, err := client.Health(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to InfluxDB at %s: %w", config.Host, err)
	}
	
	if health.Status != "pass" {
		return nil, fmt.Errorf("InfluxDB health check failed: %s", health.Message)
	}

	// Get write API for non-blocking writes
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
	bucketsAPI := sm.client.BucketsAPI()
	
	// Check if bucket exists
	bucket, err := bucketsAPI.FindBucketByName(ctx, bucketName)
	if err == nil && bucket != nil {
		log.WithField("bucket", bucketName).Info("Bucket already exists")
		return nil
	}

	// Create bucket
	_, err = bucketsAPI.CreateBucketWithName(ctx, "", bucketName, retention)
	if err != nil {
		return fmt.Errorf("failed to create bucket %s: %w", bucketName, err)
	}

	log.WithField("bucket", bucketName).Info("Bucket created successfully")
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
	
	// Close write API
	sm.writeAPI.Close()
	
	// Close client
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
