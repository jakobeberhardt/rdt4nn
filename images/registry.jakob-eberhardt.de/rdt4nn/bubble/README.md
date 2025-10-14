# Bubble Memory Stress Program

A high-performance memory interference generator designed to stress the memory subsystem with configurable working set sizes and access patterns. This program implements the bubble design principles described in academic research for creating controlled memory pressure.

## Features

- **LFSR-based Random Number Generator**: Ultra-fast pseudo-random number generation with minimal computational overhead
- **Manual SSA Memory Operations**: 100 memory operations in single static assignment form for maximum instruction-level parallelism
- **Streaming Access Patterns**: STREAM benchmark-inspired bandwidth stress testing
- **Configurable Pressure Levels**: Linear scaling from minimal to maximum memory pressure
- **OpenMP Multi-threading**: High-performance parallel execution with optimal thread management
- **Docker Support**: Easy containerized deployment
- **Real-time Configuration**: Environment variable-based configuration

## Design Principles

This bubble implementation follows key design principles for effective memory interference generation:

1. **Monotonic Curves**: Increasing pressure levels result in proportionally higher memory interference
2. **Wide Dial Range**: Pressure levels from 0.0 (minimal) to 1.0 (maximum) covering the full spectrum
3. **Broad Impact**: Comprehensive memory subsystem stress including caches, bandwidth, and prefetchers

## Architecture

### Memory Access Patterns

The bubble uses two complementary memory access patterns:

1. **Random Access Pattern**: 
   - Uses LFSR for fast random number generation
   - 100 independent memory operations in SSA form
   - Maximizes cache misses and memory latency stress

2. **Streaming Access Pattern**:
   - Based on STREAM benchmark scalar operations
   - Sequential memory access with high bandwidth utilization
   - Triggers hardware prefetchers and stresses memory bandwidth

### Working Set Calculation

Working set size is calculated based on pressure level:
```
working_set_size = min_size + (pressure_level × (max_size - min_size))
```

Where:
- `pressure_level`: 0.0 to 1.0
- `min_size`: Minimum working set size in MB
- `max_size`: Maximum working set size in MB

## Building

### Prerequisites

- GCC compiler with OpenMP support
- OpenMP runtime library
- Make (for build system)
- Docker (for containerized deployment)

### Native Build

```bash
# Build optimized version
make release

# Build debug version
make debug

# Clean build artifacts
make clean
```

### Docker Build

```bash
# Build Docker image
./build.sh build

# Or manually
docker build -t bubble .
```

## Usage

### Native Execution

```bash
# Run with default configuration
./bubble

# Run with specific pressure level
./bubble 0.8

# Run with configuration file
./run_bubble.sh

# Run with environment variables
BUBBLE_PRESSURE_LEVEL=0.7 BUBBLE_RUNTIME_SECONDS=120 ./run_bubble.sh
```

### Docker Execution

```bash
# Run with default settings
./build.sh run

# Run with specific pressure level
./build.sh run --pressure 0.8 --runtime 60

# Run quick tests
./build.sh test

# Run interactively
./build.sh run --interactive --pressure 0.5

# Run in background
./build.sh run --detach --pressure 0.9 --runtime 3600
```

## Configuration

### Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `BUBBLE_MIN_WORKING_SET_MB` | 1 | Minimum working set size in MB |
| `BUBBLE_MAX_WORKING_SET_MB` | 1024 | Maximum working set size in MB |
| `BUBBLE_NUM_THREADS` | 4 | Number of worker threads |
| `BUBBLE_RUNTIME_SECONDS` | 60 | Runtime duration (0 = infinite) |
| `BUBBLE_PRESSURE_LEVEL` | 0.5 | Pressure level (0.0-1.0) |

### Configuration File

Edit `bubble.conf` to set default values:

```bash
# Working set size range (in MB)
BUBBLE_MIN_WORKING_SET_MB=1
BUBBLE_MAX_WORKING_SET_MB=1024

# Number of worker threads
BUBBLE_NUM_THREADS=4

# Runtime duration (in seconds)
BUBBLE_RUNTIME_SECONDS=60

# Pressure level (0.0-1.0)
BUBBLE_PRESSURE_LEVEL=0.5
```

## Testing

### Quick Tests

```bash
# Test with different pressure levels
make test-low    # 0.1 pressure
make test-medium # 0.5 pressure
make test-high   # 0.9 pressure

# Docker tests
./build.sh test
```

### Performance Validation

Monitor system metrics while running bubble:

```bash
# Monitor memory bandwidth
sudo perf stat -e cache-misses,cache-references,LLC-load-misses ./bubble 0.8

# Monitor system load
htop &
./bubble 0.9

# Monitor memory usage
watch -n 1 'free -h'
```

## Use Cases

### Memory Interference Research

```bash
# Low interference baseline
BUBBLE_PRESSURE_LEVEL=0.1 ./run_bubble.sh

# Medium interference
BUBBLE_PRESSURE_LEVEL=0.5 ./run_bubble.sh

# High interference
BUBBLE_PRESSURE_LEVEL=0.9 ./run_bubble.sh
```

### Container Performance Testing

```bash
# Run bubble alongside application containers
docker run -d --name app-container my-application
docker run -d --name bubble-container -e BUBBLE_PRESSURE_LEVEL=0.6 bubble

# Monitor interference effects
docker stats
```

### Benchmark Calibration

Use bubble to create controlled interference conditions for benchmarking applications under memory pressure.

## Advanced Usage

### CPU Affinity

```bash
# Pin bubble to specific CPU cores
taskset -c 0-3 ./bubble 0.7

# Docker with CPU limits
docker run --cpuset-cpus="0-3" -e BUBBLE_PRESSURE_LEVEL=0.8 bubble
```

### Memory Limits

```bash
# Docker with memory constraints
docker run --memory="2g" -e BUBBLE_MAX_WORKING_SET_MB=1800 bubble
```

### Integration with RDT4NN

The bubble program is designed to work with the RDT4NN framework for neural network performance analysis under memory pressure:

```yaml
# Example RDT4NN configuration
containers:
  - name: workload
    image: my-neural-network:latest
  - name: bubble
    image: bubble:latest
    environment:
      - BUBBLE_PRESSURE_LEVEL=0.7
      - BUBBLE_RUNTIME_SECONDS=3600
```

## Technical Details

### LFSR Implementation

The Linear Feedback Shift Register uses mask `0xd0000001u` providing a period of 2^32 random numbers with minimal computational overhead.

### Memory Operations

Each thread performs 100 independent memory operations per iteration, ensuring maximum instruction-level parallelism and cache stress.

### Thread Design

- Even-numbered threads: Random access patterns with LFSR-based addressing
- Odd-numbered threads: Streaming access patterns based on STREAM benchmark
- OpenMP parallel regions ensure optimal thread distribution and load balancing
- Thread-local LFSR instances prevent contention and ensure independent random sequences

## Troubleshooting

### Common Issues

1. **Insufficient Memory**: Reduce `BUBBLE_MAX_WORKING_SET_MB` or number of threads
2. **Permission Denied**: Ensure executable permissions on scripts
3. **Docker Build Fails**: Check Docker daemon is running and user has permissions

### Debug Mode

```bash
# Build with debug symbols
make debug

# Run with debug output
./bubble 0.5 2>&1 | tee bubble.log
```

## License

This project is part of the RDT4NN framework. See the main project LICENSE file for details.

## Contributing

1. Follow the existing code style
2. Test with multiple pressure levels
3. Validate memory access patterns
4. Update documentation for new features

## References

Based on memory interference research principles for creating controlled memory pressure in multi-tenant environments.