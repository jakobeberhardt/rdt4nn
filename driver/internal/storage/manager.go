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
	org      string
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
	// Use the token from password field for InfluxDB 2.x authentication
	token := config.Password
	org := "rdt4nn" // Default organization for RDT4NN
	
	if config.User != "" && config.Password != "" {
		log.WithField("user", config.User).Info("InfluxDB authentication configured")
	}
	
	// Debug logging to check actual values
	log.WithFields(log.Fields{
		"host":   config.Host,
		"bucket": config.Name,
		"org":    org,
		"token":  "***", 
	}).Debug("Creating InfluxDB client with values")
	
	client := influxdb2.NewClient(config.Host, token)
	
	// Test connection with a ping (with timeout)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	health, err := client.Health(ctx)
	if err != nil {
		log.WithError(err).Warn("Failed to connect to InfluxDB - continuing without storage")
		return &Manager{
			client:   client,
			writeAPI: client.WriteAPI(org, config.Name),
			config:   config,
			org:      org,
		}, nil
	}
	
	if health.Status != "pass" {
		log.WithField("status", health.Status).Warn("InfluxDB health check warning")
	}

	writeAPI := client.WriteAPI(org, config.Name)

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
		org:      org,
	}

	log.WithFields(log.Fields{
		"host":     config.Host,
		"database": config.Name,
	}).Info("Storage manager initialized")

	return sm, nil
}

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
		
		sm.writeAPI.WritePoint(point)
	}

	log.WithField("count", len(measurements)).Debug("Measurements written to storage")
	return nil
}

func (sm *Manager) WriteMeasurement(ctx context.Context, measurement Measurement) error {
	return sm.WriteMeasurements(ctx, []Measurement{measurement})
}

func (sm *Manager) Query(ctx context.Context, query string) ([]map[string]interface{}, error) {
	queryAPI := sm.client.QueryAPI(sm.org)
	
	result, err := queryAPI.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var results []map[string]interface{}
	for result.Next() {
		record := result.Record()
		recordMap := make(map[string]interface{})
		
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
	
	log.WithField("bucket", bucketName).Info("Bucket creation skipped - ensure bucket exists in InfluxDB")
	return nil
}

func (sm *Manager) Flush() error {
	sm.writeAPI.Flush()
	return nil
}

func (sm *Manager) Close() error {
	log.Info("Closing storage manager")
	
	sm.writeAPI.Flush()
	
	sm.client.Close()
	
	return nil
}

func (sm *Manager) GetStats(ctx context.Context) (map[string]interface{}, error) {
	stats := make(map[string]interface{})
	
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
