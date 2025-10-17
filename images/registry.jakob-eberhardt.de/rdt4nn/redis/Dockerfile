FROM redis:7-alpine

RUN apk add --no-cache python3 py3-pip && \
    pip3 install redis --break-system-packages

COPY generate_data.py /tmp/generate_data.py

RUN mkdir -p /data

# This bakes the data into the image
RUN redis-server --daemonize yes --dir /data && \
    sleep 2 && \
    python3 /tmp/generate_data.py && \
    redis-cli SAVE && \
    redis-cli SHUTDOWN NOSAVE || true && \
    sleep 2 && \
    rm /tmp/generate_data.py

EXPOSE 6379

CMD ["redis-server", "--dir", "/data", "--appendonly", "no"]