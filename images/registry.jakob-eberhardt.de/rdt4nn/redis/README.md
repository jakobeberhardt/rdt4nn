# Redis Prepopulated Image (64-byte keys)

This directory contains a Dockerfile and scripts to create a Redis image prepopulated with approximately 1GB of data for benchmarking and sensitivity testing.

## Image Characteristics

- **Base Image**: redis:7-alpine
- **Key Size**: 64 bytes (format: `key_XXXXXXX...` with random alphanumeric suffix)
- **Value Size**: Variable, 512-2048 bytes per value (average ~1KB)
- **Total Data**: ~1GB in Redis memory
- **Estimated Keys**: ~700,000-900,000 keys (depending on actual value sizes)

## Data Structure

The data is generated with the following properties:

- **Keys**: Fixed 64-byte strings starting with `key_` followed by random alphanumeric characters
- **Values**: Random alphanumeric strings with variable length (512-2048 bytes)
- **Distribution**: Uniformly distributed across the keyspace

## Building the Image

```bash
docker build -t redis-prepopulated-64:latest .
```


```bash
docker run -d -p 6379:6379 --name redis-bench redis-prepopulated-64:latest
```

### Verifying the Data

Check the number of keys:
```bash
docker exec redis-bench redis-cli DBSIZE
```

Check memory usage:
```bash
docker exec redis-bench redis-cli INFO memory | grep used_memory_human
```

Sample some keys:
```bash
docker exec redis-bench redis-cli RANDOMKEY
docker exec redis-bench redis-cli GET <key_name>
```

### Using in Container-Bench

Example YAML configuration:

```yaml
benchmark:
  name: Redis Sensitivity Test
  description: Test Redis performance under noisy neighbor scenarios
  max_t: 60
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

container0:
  index: 0
  name: redis_workload
  image: registry.jakob-eberhardt.de/rdt4nn/redis-prepopulated-64:latest
  core: 0
  port: 6379:6379
  data:
    frequency: 100
    perf: true
    docker: true
    rdt: true
```

## Pushing to Registry

After building, push to your registry:

```bash
docker tag redis-prepopulated-64:latest registry.jakob-eberhardt.de/rdt4nn/redis-prepopulated-64:latest
docker push registry.jakob-eberhardt.de/rdt4nn/redis-prepopulated-64:latest
```
