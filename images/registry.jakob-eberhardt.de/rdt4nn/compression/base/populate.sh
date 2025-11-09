#!/bin/sh

set -e

if [ -z "$1" ]; then
    echo "Usage: $0 <size_in_mb> [output_file]"
    echo "Example: $0 5000 /data/data.bin"
    exit 1
fi

SIZE_MB=$1
OUTPUT_FILE=${2:-/data/data.bin}

echo "Generating ${SIZE_MB}MB file at ${OUTPUT_FILE}..."
dd if=/dev/urandom of="${OUTPUT_FILE}" bs=1M count="${SIZE_MB}" 2>&1 | tail -n 3
chmod 644 "${OUTPUT_FILE}"
echo "File created successfully: ${OUTPUT_FILE} (${SIZE_MB}MB)"
ls -lh "${OUTPUT_FILE}"
