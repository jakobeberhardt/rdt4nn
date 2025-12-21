FROM jrottenberg/ffmpeg:latest
RUN apt-get update && apt-get install -y bc wget && rm -rf /var/lib/apt/lists/*
RUN mkdir /data
COPY generate_video.sh /generate_video.sh
RUN chmod +x /generate_video.sh
WORKDIR /data
ENTRYPOINT []
CMD ["sleep", "infinity"] 