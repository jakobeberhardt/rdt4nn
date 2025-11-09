FROM ubuntu:22.04
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update && apt-get install -y \
    build-essential \
    gfortran \
    libopenmpi-dev \
    openmpi-bin \
    git \
    vim \
    && rm -rf /var/lib/apt/lists/*

RUN git clone --recursive https://github.com/UK-MAC/TeaLeaf.git /TeaLeaf
WORKDIR /TeaLeaf/TeaLeaf_ref
RUN make COMPILER=GNU OPTIONS="-fallow-argument-mismatch"
CMD ["sleep", "infinity"] 