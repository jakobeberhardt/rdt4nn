FROM gcc:12-bullseye AS builder

WORKDIR /build

COPY bubble.c .

RUN gcc -O3 -fopenmp -o bubble bubble.c -lm

FROM debian:bullseye-slim

RUN apt-get update && \
    apt-get install -y libgomp1 tini && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /build/bubble /usr/local/bin/bubble

ENTRYPOINT ["/usr/local/bin/bubble"]
