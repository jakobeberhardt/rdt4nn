package cmd

import (
	"fmt"
	"os"

	"github.com/jakobeberhardt/rdt4nn/driver/internal/benchmark"
	"github.com/jakobeberhardt/rdt4nn/driver/internal/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	configFile string
	verbose    bool
)

var rootCmd = &cobra.Command{
	Use:   "rdt4nn-driver",
	Short: "A comprehensive container benchmarking tool for Intel RDT performance analysis",
	Long: `rdt4nn-driver is a tool for defining, running, and profiling container-based benchmarks
on Intel RDT-enabled machines to study the busy neighbor problem regarding memory bandwidth and caches.

The tool provides:
- YAML-based benchmark configuration
- Container orchestration using Docker
- Hardware-level profiling using Intel RDT and perf
- Data collection and storage in InfluxDB
- Pluggable scheduler implementations`,
	Run: runBenchmark,
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "Path to benchmark configuration file (required)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose logging")
	rootCmd.MarkPersistentFlagRequired("config")
}

func Execute() error {
	return rootCmd.Execute()
}

func runBenchmark(cmd *cobra.Command, args []string) {
	// Set up logging
	if verbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	log.Info("Starting RDT4NN benchmark driver")

	// Load configuration
	cfg, err := config.LoadConfig(configFile)
	if err != nil {
		log.WithError(err).Fatal("Failed to load configuration")
		os.Exit(1)
	}

	// Set log level from config
	logLevel, err := log.ParseLevel(cfg.Benchmark.LogLevel)
	if err != nil {
		log.WithError(err).Warn("Invalid log level in config, using INFO")
		logLevel = log.InfoLevel
	}
	log.SetLevel(logLevel)

	log.WithField("config", configFile).Info("Configuration loaded successfully")

	// Create and run benchmark
	benchmarkRunner, err := benchmark.NewRunner(cfg)
	if err != nil {
		log.WithError(err).Fatal("Failed to create benchmark runner")
		os.Exit(1)
	}

	if err := benchmarkRunner.Run(); err != nil {
		log.WithError(err).Fatal("Benchmark execution failed")
		os.Exit(1)
	}

	log.Info("Benchmark completed successfully")
}

func init() {
	// Add version command
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of rdt4nn-driver",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("rdt4nn-driver v1.0.0")
		},
	}
	rootCmd.AddCommand(versionCmd)

	// Add validate command
	var validateCmd = &cobra.Command{
		Use:   "validate",
		Short: "Validate benchmark configuration file",
		Run: func(cmd *cobra.Command, args []string) {
			if configFile == "" {
				log.Fatal("Configuration file is required")
			}
			
			cfg, err := config.LoadConfig(configFile)
			if err != nil {
				log.WithError(err).Fatal("Configuration validation failed")
			}
			
			log.WithField("benchmark", cfg.Benchmark.Name).Info("Configuration is valid")
		},
	}
	rootCmd.AddCommand(validateCmd)
}
