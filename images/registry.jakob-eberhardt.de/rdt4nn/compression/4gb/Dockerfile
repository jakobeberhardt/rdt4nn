FROM alpine:latest

RUN apk add --no-cache \
    gzip \
    bzip2 \
    xz \
    zstd \
    lz4
RUN mkdir /data
RUN dd if=/dev/urandom of=/data/data1.bin bs=1M count=4096 && \
    chmod 644 /data/data1.bin

WORKDIR /data

CMD ["sleep", "infinity"]
