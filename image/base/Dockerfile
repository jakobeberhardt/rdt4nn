FROM ubuntu:22.04
ARG DEBIAN_FRONTEND=noninteractive
RUN apt-get update && apt upgrade
RUN echo "This image is part of rdt4nn (https://github.com/jakobeberhardt/rdt4nn).\nBuild: $(date)" > /release.txt