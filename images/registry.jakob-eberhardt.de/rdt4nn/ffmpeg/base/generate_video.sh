#!/bin/sh

set -e

if [ -z "$1" ]; then
    echo "Usage: $0 <size_in_mb> [output_file]"
    echo "Example: $0 20 /data/video.mov"
    exit 1
fi

SIZE_MB=$1
OUTPUT_FILE=${2:-/data/video.mov}

# Calculate duration based on target size
BITRATE=2
DURATION=$(echo "scale=0; ($SIZE_MB * 8) / $BITRATE" | bc)

echo "Generating ${SIZE_MB}MB video at ${OUTPUT_FILE}..."
echo "Duration: ${DURATION} seconds, Video bitrate: ${BITRATE}Mbps"

ffmpeg -f lavfi -i testsrc=duration=${DURATION}:size=1920x1080:rate=30 \
       -f lavfi -i sine=frequency=1000:duration=${DURATION} \
       -c:v libx264 -preset fast -b:v ${BITRATE}M \
       -c:a aac -b:a 128k \
       -movflags +faststart \
       "${OUTPUT_FILE}" -y 2>&1 | tail -n 10

chmod 644 "${OUTPUT_FILE}"
ACTUAL_SIZE=$(du -h "${OUTPUT_FILE}" | cut -f1)
echo "Video created successfully: ${OUTPUT_FILE}"
echo "Target size: ${SIZE_MB}MB, Actual size: ${ACTUAL_SIZE}"
ls -lh "${OUTPUT_FILE}"
