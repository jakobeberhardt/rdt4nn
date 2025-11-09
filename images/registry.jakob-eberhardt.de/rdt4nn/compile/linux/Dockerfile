FROM ubuntu:22.04

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -y \
    build-essential \
    bc \
    bison \
    flex \
    libssl-dev \
    libelf-dev \
    libncurses-dev \
    git \
    wget \
    curl \
    kmod \
    cpio \
    dwarves \
    rsync \
    vim \
    && rm -rf /var/lib/apt/lists/*

RUN git clone --depth=1 https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git /linux

WORKDIR /linux

CMD ["sleep", "infinity"]
