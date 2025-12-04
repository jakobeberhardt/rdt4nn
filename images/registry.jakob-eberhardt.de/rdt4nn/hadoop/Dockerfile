FROM apache/hadoop:3.4.1

#USER root

RUN mkdir -p /opt/hadoop/etc/hadoop

COPY core-site.xml /opt/hadoop/etc/hadoop/core-site.xml
COPY hdfs-site.xml /opt/hadoop/etc/hadoop/hdfs-site.xml
COPY benchmark.sh /benchmark.sh

RUN sudo chmod +x /benchmark.sh

ENV BENCHMARK_TYPE=mixed
ENV DATA_SIZE=medium
ENV ITERATIONS=10
ENV WARMUP_TIME=20

USER hadoop

CMD ["/benchmark.sh"]
