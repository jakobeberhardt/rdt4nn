#!/bin/bash
set -e

echo "================================"
echo "Hadoop Benchmark Configuration"
echo "================================"
echo "Benchmark Type: $BENCHMARK_TYPE"
echo "Data Size: $DATA_SIZE"
echo "Iterations: $ITERATIONS"
echo "Warmup Time: $WARMUP_TIME seconds"
echo "================================"

# Unset niceness to avoid priority errors
unset HADOOP_NICENESS

# Initialize HDFS
echo "Formatting HDFS namenode..."
hdfs namenode -format -force

echo "Starting HDFS namenode..."
hdfs --daemon start namenode
sleep 5

echo "Starting HDFS datanode..."
hdfs --daemon start datanode
sleep "$WARMUP_TIME"

echo "HDFS services started successfully"
hdfs dfsadmin -report

# Parse DATA_SIZE environment variable (e.g., "2GB", "7GB", "15GB")
if [[ "$DATA_SIZE" =~ ^([0-9]+)GB$ ]]; then
    SIZE_GB="${BASH_REMATCH[1]}"
    DATA_SIZE_BYTES=$((SIZE_GB * 1024 * 1024 * 1024))
else
    echo "Invalid DATA_SIZE format: $DATA_SIZE. Expected format: <number>GB (e.g., 7GB, 12GB)"
    echo "Using default: 7GB"
    DATA_SIZE_BYTES=7516192768      # 7GB
fi

# Use MAPS_COUNT from environment variable, default to 2
if [ -z "$MAPS_COUNT" ]; then
    MAPS_COUNT=2
fi

echo "Generating benchmark dataset..."
echo "  - Dataset size: $(($DATA_SIZE_BYTES/1024/1024))MB with $MAPS_COUNT maps"
echo "New Script!"

# Determine which data generator to use based on BENCHMARK_TYPE
if [ "$BENCHMARK_TYPE" = "sort" ]; then
    # TeraSort requires TeraGen data
    RECORD_COUNT=$((DATA_SIZE_BYTES / 100))  # Each record is 100 bytes
    echo "  - Using TeraGen for sort benchmark: $RECORD_COUNT records"
    hadoop jar /opt/hadoop/share/hadoop/mapreduce/hadoop-mapreduce-examples-*.jar teragen \
        -Dmapreduce.job.maps=$MAPS_COUNT \
        $RECORD_COUNT \
        hdfs://localhost:9000/benchmark-data
else
    # Other benchmarks use RandomWriter
    echo "  - Using RandomWriter for ${BENCHMARK_TYPE} benchmark"
    hadoop jar /opt/hadoop/share/hadoop/mapreduce/hadoop-mapreduce-examples-*.jar randomwriter \
        -Dmapreduce.job.maps=$MAPS_COUNT \
        -Dtest.randomwrite.total_bytes=$DATA_SIZE_BYTES \
        hdfs://localhost:9000/benchmark-data
fi

echo "Data generation completed!"

# Run benchmark based on BENCHMARK_TYPE
case "$BENCHMARK_TYPE" in
    wordcount)
        if [ "$ITERATIONS" -eq -1 ]; then
            echo "Running WordCount benchmark indefinitely..."
            i=1
            while true; do
                echo "=== WordCount Iteration $i ==="
                hdfs dfs -rm -r -f hdfs://localhost:9000/output-wordcount
                hadoop jar /opt/hadoop/share/hadoop/mapreduce/hadoop-mapreduce-examples-*.jar wordcount \
                    hdfs://localhost:9000/benchmark-data \
                    hdfs://localhost:9000/output-wordcount
                i=$((i+1))
            done
        else
            echo "Running WordCount benchmark for $ITERATIONS iterations..."
            for i in $(seq 1 $ITERATIONS); do
                echo "=== WordCount Iteration $i/$ITERATIONS ==="
                hdfs dfs -rm -r -f hdfs://localhost:9000/output-wordcount
                hadoop jar /opt/hadoop/share/hadoop/mapreduce/hadoop-mapreduce-examples-*.jar wordcount \
                    hdfs://localhost:9000/benchmark-data \
                    hdfs://localhost:9000/output-wordcount
            done
        fi
        ;;
    
    grep)
        if [ "$ITERATIONS" -eq -1 ]; then
            echo "Running Grep benchmark indefinitely..."
            i=1
            while true; do
                echo "=== Grep Iteration $i ==="
                hdfs dfs -rm -r -f hdfs://localhost:9000/output-grep
                hadoop jar /opt/hadoop/share/hadoop/mapreduce/hadoop-mapreduce-examples-*.jar grep \
                    hdfs://localhost:9000/benchmark-data \
                    hdfs://localhost:9000/output-grep \
                    'dfs[a-z.]+'
                i=$((i+1))
            done
        else
            echo "Running Grep benchmark for $ITERATIONS iterations..."
            for i in $(seq 1 $ITERATIONS); do
                echo "=== Grep Iteration $i/$ITERATIONS ==="
                hdfs dfs -rm -r -f hdfs://localhost:9000/output-grep
                hadoop jar /opt/hadoop/share/hadoop/mapreduce/hadoop-mapreduce-examples-*.jar grep \
                    hdfs://localhost:9000/benchmark-data \
                    hdfs://localhost:9000/output-grep \
                    'dfs[a-z.]+'
            done
        fi
        ;;
    
    sort)
        if [ "$ITERATIONS" -eq -1 ]; then
            echo "Running Sort benchmark indefinitely..."
            i=1
            while true; do
                echo "=== Sort Iteration $i ==="
                hdfs dfs -rm -r -f hdfs://localhost:9000/output-sort
                hadoop jar /opt/hadoop/share/hadoop/mapreduce/hadoop-mapreduce-examples-*.jar terasort \
                    hdfs://localhost:9000/benchmark-data \
                    hdfs://localhost:9000/output-sort
                i=$((i+1))
            done
        else
            echo "Running Sort benchmark for $ITERATIONS iterations..."
            for i in $(seq 1 $ITERATIONS); do
                echo "=== Sort Iteration $i/$ITERATIONS ==="
                hdfs dfs -rm -r -f hdfs://localhost:9000/output-sort
                hadoop jar /opt/hadoop/share/hadoop/mapreduce/hadoop-mapreduce-examples-*.jar terasort \
                    hdfs://localhost:9000/benchmark-data \
                    hdfs://localhost:9000/output-sort
            done
        fi
        ;;
    
    mixed)
        if [ "$ITERATIONS" -eq -1 ]; then
            echo "Running Mixed benchmark indefinitely..."
            i=1
            while true; do
                echo "=== Mixed Iteration $i ==="
                
                echo "  Running WordCount..."
                hdfs dfs -rm -r -f hdfs://localhost:9000/output-wordcount
                hadoop jar /opt/hadoop/share/hadoop/mapreduce/hadoop-mapreduce-examples-*.jar wordcount \
                    hdfs://localhost:9000/benchmark-data \
                    hdfs://localhost:9000/output-wordcount
                
                echo "  Running Grep..."
                hdfs dfs -rm -r -f hdfs://localhost:9000/output-grep
                hadoop jar /opt/hadoop/share/hadoop/mapreduce/hadoop-mapreduce-examples-*.jar grep \
                    hdfs://localhost:9000/benchmark-data \
                    hdfs://localhost:9000/output-grep \
                    'dfs[a-z.]+'
                
                echo "  Running Pi estimation..."
                hadoop jar /opt/hadoop/share/hadoop/mapreduce/hadoop-mapreduce-examples-*.jar pi \
                    $MAPS_COUNT 1000
                
                i=$((i+1))
            done
        else
            echo "Running Mixed benchmark for $ITERATIONS iterations..."
            for i in $(seq 1 $ITERATIONS); do
                echo "=== Mixed Iteration $i/$ITERATIONS ==="
                
                echo "  Running WordCount..."
                hdfs dfs -rm -r -f hdfs://localhost:9000/output-wordcount
                hadoop jar /opt/hadoop/share/hadoop/mapreduce/hadoop-mapreduce-examples-*.jar wordcount \
                    hdfs://localhost:9000/benchmark-data \
                    hdfs://localhost:9000/output-wordcount
                
                echo "  Running Grep..."
                hdfs dfs -rm -r -f hdfs://localhost:9000/output-grep
                hadoop jar /opt/hadoop/share/hadoop/mapreduce/hadoop-mapreduce-examples-*.jar grep \
                    hdfs://localhost:9000/benchmark-data \
                    hdfs://localhost:9000/output-grep \
                    'dfs[a-z.]+'
                
                echo "  Running Pi estimation..."
                hadoop jar /opt/hadoop/share/hadoop/mapreduce/hadoop-mapreduce-examples-*.jar pi \
                    $MAPS_COUNT 1000
                
            done
        fi
        ;;
    
    *)
        echo "Unknown BENCHMARK_TYPE: $BENCHMARK_TYPE"
        echo "Valid types: wordcount, grep, sort, mixed"
        exit 1
        ;;
esac

echo "================================"
echo "Benchmark completed successfully!"
echo "================================"
hdfs dfs -du -h /
