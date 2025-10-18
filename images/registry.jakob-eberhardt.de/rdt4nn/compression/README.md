# gzip compression
cat /data/*.bin | gzip -c > /dev/null

# bzip2 compression (more CPU intensive)
cat /data/*.bin | bzip2 -c > /dev/null

# xz compression (most CPU intensive, best compression)
cat /data/*.bin | xz -c > /dev/null

# zstd compression (modern, fast)
cat /data/*.bin | zstd -c > /dev/null

# lz4 compression (fastest)
cat /data/*.bin | lz4 -c > /dev/null

# Continuous loop for benchmarking
while true; do cat /data/*.bin | gzip -c > /dev/null; done