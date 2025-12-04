# Hadoop Benchmark Docker Image

A pre-configured Hadoop container for running standardized benchmarks with HDFS.

## Features

- Based on `apache/hadoop:3.4.1`
- Pre-configured HDFS (namenode + datanode)
- Configurable benchmark types and data sizes via environment variables
- Automatic data generation and benchmark execution
- Web UI accessible on port 9870


## Environment Variables

| Variable | Options | Default | Description |
|----------|---------|---------|-------------|
| `BENCHMARK_TYPE` | wordcount, grep, sort, mixed | mixed | Type of benchmark to run |
| `DATA_SIZE` | small, medium, large | medium | Size of test datasets |
| `ITERATIONS` | any number or -1 | 10 | Number of benchmark iterations (-1 = run indefinitely) |
| `WARMUP_TIME` | seconds | 20 | Time to wait for HDFS startup |

## Data Sizes

### Small
- Large dataset: 1GB (4 maps)
- Medium dataset: 512MB (2 maps)
- Small dataset: 256MB (2 maps)

### Medium
- Large dataset: 2GB (8 maps)
- Medium dataset: 1GB (4 maps)
- Small dataset: 512MB (2 maps)

### Large
- Large dataset: 5GB (16 maps)
- Medium dataset: 2GB (8 maps)
- Small dataset: 1GB (4 maps)

## Benchmark Types

### wordcount
Runs MapReduce word count on the large dataset

### grep
Runs pattern matching (grep) on the medium dataset

### sort
Runs TeraSort on the medium dataset

### mixed
Runs a combination of WordCount, Grep, and Pi estimation

## Usage Examples

### Default (mixed benchmark, medium data, 10 iterations)
```bash
docker run -it --rm --privileged -p 9870:9870 hadoop-bench:latest
```

### Run indefinitely (until container is stopped)
```bash
docker run -it --rm --privileged -p 9870:9870 \
  -e BENCHMARK_TYPE=wordcount \
  -e DATA_SIZE=medium \
  -e ITERATIONS=-1 \
  hadoop-bench:latest
```

### Access web UI
Open http://localhost:9870 in your browser while the container is running.

## Use with container-bench

Example YAML configuration:

```yaml
benchmark:
  name: Hadoop Benchmark
  description: 10-minute Hadoop mixed workload
  max_t: 630
  log_level: info
  scheduler:
    implementation: default
    rdt: false
  data:
    db:
      host: ${INFLUXDB_HOST}
      name: ${INFLUXDB_BUCKET}
      user: ${INFLUXDB_USER}
      password: ${INFLUXDB_TOKEN}
      org: ${INFLUXDB_ORG}

hadoop:
  index: 0
  image: hadoop-bench:latest
  port: 9870:9870
  core: 0
  privileged: true
  environment:
    - BENCHMARK_TYPE=mixed
    - DATA_SIZE=medium
    - ITERATIONS=10
    - WARMUP_TIME=20
  data:
    frequency: 100
    perf:
      instructions: true
      cycles: true
      cache_references: true
      cache_misses: true
    docker:
      cpu_usage_percent: true
      memory_usage: true
    rdt: false
```
