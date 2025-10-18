FROM registry.jakob-eberhardt.de/rdt4nn/compression/4gb:latest

RUN dd if=/dev/urandom of=/data/data2.bin bs=1M count=6144 && \
    chmod 644 /data/data2.bin

WORKDIR /data

CMD ["sleep", "infinity"]
