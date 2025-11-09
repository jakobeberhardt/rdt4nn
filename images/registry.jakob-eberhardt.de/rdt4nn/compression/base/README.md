# Compression Base Image

Base image for compression benchmarks with pre-installed compression tools and a data population script.

## Installed Tools

- gzip
- bzip2
- xz
- zstd
- lz4
- 7zip

## populate.sh Script

The image includes `/populate.sh` for generating random data files on demand.

### Usage

```bash
/populate.sh <size_in_mb> [output_file]
```

### Examples

Generate a 5GB file:
```bash
/populate.sh 5000
```

Generate a 1GB file with custom name:
```bash
/populate.sh 1024 /data/mydata.bin
```

## Building

```bash
docker build -t registry.jakob-eberhardt.de/rdt4nn/compression/base:latest .
```

## Running

Run interactively and populate on demand:
```bash
docker run -it --rm registry.jakob-eberhardt.de/rdt4nn/compression/base:latest sh
# Inside container:
/populate.sh 5000
```

## Derived Images

Derived images can extend this base and pre-populate data during build:

```dockerfile
FROM registry.jakob-eberhardt.de/rdt4nn/compression/base:latest

RUN /populate.sh 4096 /data/data1.bin

CMD ["sleep", "infinity"]
```
