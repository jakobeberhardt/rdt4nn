FROM ubuntu:22.04
RUN apt-get update && apt-get install -y \
    gcc \
    make \
    && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY bubble.c .
COPY bubble.conf .
COPY run_bubble.sh .
COPY Makefile .
RUN make release
RUN useradd -m -s /bin/bash bubble
USER bubble
ENV BUBBLE_MIN_WORKING_SET_MB=1
ENV BUBBLE_MAX_WORKING_SET_MB=1024
ENV BUBBLE_NUM_THREADS=4
ENV BUBBLE_RUNTIME_SECONDS=60
ENV BUBBLE_PRESSURE_LEVEL=0.5
CMD ["./run_bubble.sh"]