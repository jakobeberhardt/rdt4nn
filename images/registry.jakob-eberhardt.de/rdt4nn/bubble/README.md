# Bubble - Memory Pressure Generator

A containerized memory pressure generator for studying container interference in shared resource environments. Based on the bubble design pattern for generating controlled memory subsystem pressure.

## Overview

The bubble application creates configurable memory pressure through:
- **LFSR-based random access patterns** - Minimizes computation overhead between memory operations
- **Manual SSA (Single Static Assignment)** - 100 independent memory operations for high ILP
- **STREAM-based bandwidth stressing** - Triggers prefetchers and stresses memory bandwidth

## Quick Start

### Build

```bash
docker build . -t registry.jakob-eberhardt.de/rdt4nn/bubble:latest
```

### Run with Defaults

```bash
docker run --rm registry.jakob-eberhardt.de/rdt4nn/bubble
```

Default configuration: 100 MB working set, all available CPU cores, 50% streaming/random mix.

## Configuration

### Command-Line Options

| Option | Description | Default |
|--------|-------------|---------|
| `--min-size MB` | Minimum working set size | 1 MB |
| `--max-size MB` | Maximum working set size | 100 MB |
| `--ramp-time SECONDS` | Ramp duration from min to max | 0 (instant) |
| `--threads N` | Number of threads | System cores |
| `--streaming-ratio %` | Percentage doing streaming (vs random) | 50% |

### Examples

**Simple pressure test (1 GB working set):**
```bash
docker run --rm registry.jakob-eberhardt.de/rdt4nn/bubble \
  --max-size 1024
```

**Gradual pressure increase over 60 seconds:**
```bash
docker run --rm registry.jakob-eberhardt.de/rdt4nn/bubble \
  --min-size 10 \
  --max-size 1000 \
  --ramp-time 60
```

**Fixed thread count with more random access:**
```bash
docker run --rm registry.jakob-eberhardt.de/rdt4nn/bubble \
  --threads 4 \
  --streaming-ratio 25 \
  --max-size 512
```



## Design Principles

### Monotonic Pressure
As working set size increases, memory subsystem pressure increases monotonically. Higher pressure → more interference.

### Wide Dial Range
- Minimum: Near-zero pressure (1 MB)
- Maximum: Configurable up to available memory
- Gradual ramping supported for controlled studies

### Broad Impact
- **Random access** stresses cache hierarchy and random access latency
- **Streaming access** stresses memory bandwidth and prefetchers
- **Configurable mix** allows targeting specific interference patterns

## Technical Details

### LFSR Implementation
Uses mask `0xd0000001u` with period of 2^32, providing fast pseudo-random generation with minimal computational overhead.

### Manual SSA Block
100 independent memory operations per iteration eliminate data dependencies, maximizing instruction-level parallelism and memory subsystem utilization.

### OpenMP Threading
Each thread operates on shared data structures with thread-local LFSR state, creating realistic multi-core interference patterns.

## Integration with RDT4NN

The bubble can be integrated into RDT4NN experiment configurations: