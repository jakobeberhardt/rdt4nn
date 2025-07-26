FROM registry.jakob-eberhardt.de/rdt4nn/base/ubuntu:22.04.0 AS builder
ARG DEBIAN_FRONTEND=noninteractive
RUN apt-get update && \
    apt-get install -y --no-install-recommends build-essential && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /opt/matmul
COPY matmul.cpp .
RUN g++ -O3 -march=native -std=c++17 -o matmul matmul.cpp

FROM registry.jakob-eberhardt.de/rdt4nn/base/ubuntu:22.04.0
ARG DEBIAN_FRONTEND=noninteractive
RUN apt-get update && \
    apt-get install -y --no-install-recommends libstdc++6 && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /opt/matmul/matmul /usr/local/bin/

WORKDIR /opt/matmul
ENTRYPOINT ["/usr/local/bin/matmul", "input.in"]
