FROM registry.jakob-eberhardt.de/rdt4nn/base/ubuntu:22.04.0 AS builder
ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        build-essential \
        gfortran \
        git \
        make \
        mpich \
        libmpich-dev && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /opt
ENV GIT_SSL_NO_VERIFY=1
RUN git clone --depth 1 https://github.com/UK-MAC/TeaLeaf_ref.git

WORKDIR /opt/TeaLeaf_ref
RUN make -j$(nproc) COMPILER=GNU

FROM registry.jakob-eberhardt.de/rdt4nn/base/ubuntu:22.04.0 AS tealeaf
ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        mpich \
        libgfortran5 \
		libgomp1 && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /opt/TeaLeaf_ref/tea_leaf        /usr/local/bin/
COPY --from=builder /opt/TeaLeaf_ref/tea.in          /opt/tealeaf/

WORKDIR /opt/tealeaf
ENTRYPOINT ["mpirun", "-np", "1", "/usr/local/bin/tea_leaf"]