FROM alpine:latest

RUN apk add --no-cache \
    gzip \
    bzip2 \
    xz \
    zstd \
    lz4 \
    7zip

RUN mkdir /data

COPY populate.sh /populate.sh
RUN chmod +x /populate.sh

WORKDIR /data

CMD ["sleep", "infinity"]
