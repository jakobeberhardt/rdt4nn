# FFmpeg Base Image

Base image for video encoding/transcoding benchmarks with ffmpeg and a video generation script.

## Features

- Based on jrottenberg/ffmpeg (includes full ffmpeg with all codecs)
- Video generation script for creating test videos of specific sizes
- Supports H.264 video with AAC audio encoding

## generate_video.sh Script

The image includes `/generate_video.sh` for generating test videos on demand.

### Usage

```bash
/generate_video.sh <size_in_mb> [output_file]
```

### Examples

Generate a 20MB video:
```bash
/generate_video.sh 20
```

Generate a 500MB video with custom name:
```bash
/generate_video.sh 500 /data/large_video.mov
```

## Building

```bash
docker build -t registry.jakob-eberhardt.de/rdt4nn/ffmpeg/base:latest .
```

## Running

Run interactively and generate video on demand:
```bash
docker run -it --rm registry.jakob-eberhardt.de/rdt4nn/ffmpeg/base:latest sh
# Inside container:
/generate_video.sh 100
```