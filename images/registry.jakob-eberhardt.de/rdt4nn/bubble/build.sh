#!/bin/bash

# Bubble Docker Build and Run Script

IMAGE_NAME="bubble"
TAG="latest"
FULL_IMAGE_NAME="${IMAGE_NAME}:${TAG}"

# Default configuration
DEFAULT_PRESSURE="0.5"
DEFAULT_RUNTIME="60"
DEFAULT_THREADS="4"
DEFAULT_MIN_WS="1"
DEFAULT_MAX_WS="1024"

# Function to show usage
show_usage() {
    echo "Usage: $0 [COMMAND] [OPTIONS]"
    echo ""
    echo "Commands:"
    echo "  build                 Build the Docker image"
    echo "  run [OPTIONS]         Run the bubble container"
    echo "  test                  Run quick tests with different pressure levels"
    echo "  clean                 Remove the Docker image"
    echo ""
    echo "Run Options:"
    echo "  -p, --pressure LEVEL  Set pressure level (0.0-1.0, default: $DEFAULT_PRESSURE)"
    echo "  -t, --runtime SECONDS Set runtime in seconds (default: $DEFAULT_RUNTIME)"
    echo "  -n, --threads COUNT   Set number of threads (default: $DEFAULT_THREADS)"
    echo "  -m, --min-ws SIZE     Set minimum working set size in MB (default: $DEFAULT_MIN_WS)"
    echo "  -M, --max-ws SIZE     Set maximum working set size in MB (default: $DEFAULT_MAX_WS)"
    echo "  -d, --detach          Run container in background"
    echo "  -i, --interactive     Run container interactively"
    echo ""
    echo "Examples:"
    echo "  $0 build"
    echo "  $0 run -p 0.8 -t 30"
    echo "  $0 run --pressure 0.2 --runtime 120 --threads 8"
    echo "  $0 test"
}

# Function to build Docker image
build_image() {
    echo "Building Docker image: $FULL_IMAGE_NAME"
    docker build -t "$FULL_IMAGE_NAME" .
    if [ $? -eq 0 ]; then
        echo "Successfully built $FULL_IMAGE_NAME"
    else
        echo "Failed to build Docker image"
        exit 1
    fi
}

# Function to run container
run_container() {
    local pressure="$DEFAULT_PRESSURE"
    local runtime="$DEFAULT_RUNTIME"
    local threads="$DEFAULT_THREADS"
    local min_ws="$DEFAULT_MIN_WS"
    local max_ws="$DEFAULT_MAX_WS"
    local docker_args=""
    local detach=false
    local interactive=false
    
    # Parse arguments
    while [[ $# -gt 0 ]]; do
        case $1 in
            -p|--pressure)
                pressure="$2"
                shift 2
                ;;
            -t|--runtime)
                runtime="$2"
                shift 2
                ;;
            -n|--threads)
                threads="$2"
                shift 2
                ;;
            -m|--min-ws)
                min_ws="$2"
                shift 2
                ;;
            -M|--max-ws)
                max_ws="$2"
                shift 2
                ;;
            -d|--detach)
                detach=true
                shift
                ;;
            -i|--interactive)
                interactive=true
                shift
                ;;
            *)
                echo "Unknown option: $1"
                show_usage
                exit 1
                ;;
        esac
    done
    
    # Set Docker run arguments
    if [ "$detach" = true ]; then
        docker_args="$docker_args -d"
    fi
    
    if [ "$interactive" = true ]; then
        docker_args="$docker_args -it"
    else
        docker_args="$docker_args --rm"
    fi
    
    echo "Running bubble container with:"
    echo "  Pressure Level: $pressure"
    echo "  Runtime: $runtime seconds"
    echo "  Threads: $threads"
    echo "  Working Set Range: $min_ws - $max_ws MB"
    
    docker run $docker_args \
        -e "BUBBLE_PRESSURE_LEVEL=$pressure" \
        -e "BUBBLE_RUNTIME_SECONDS=$runtime" \
        -e "BUBBLE_NUM_THREADS=$threads" \
        -e "BUBBLE_MIN_WORKING_SET_MB=$min_ws" \
        -e "BUBBLE_MAX_WORKING_SET_MB=$max_ws" \
        "$FULL_IMAGE_NAME"
}

# Function to run tests
run_tests() {
    echo "Running bubble tests with different pressure levels..."
    
    echo "=== Test 1: Low Pressure (0.1) ==="
    docker run --rm \
        -e "BUBBLE_PRESSURE_LEVEL=0.1" \
        -e "BUBBLE_RUNTIME_SECONDS=10" \
        -e "BUBBLE_NUM_THREADS=2" \
        "$FULL_IMAGE_NAME"
    
    echo ""
    echo "=== Test 2: Medium Pressure (0.5) ==="
    docker run --rm \
        -e "BUBBLE_PRESSURE_LEVEL=0.5" \
        -e "BUBBLE_RUNTIME_SECONDS=10" \
        -e "BUBBLE_NUM_THREADS=4" \
        "$FULL_IMAGE_NAME"
    
    echo ""
    echo "=== Test 3: High Pressure (0.9) ==="
    docker run --rm \
        -e "BUBBLE_PRESSURE_LEVEL=0.9" \
        -e "BUBBLE_RUNTIME_SECONDS=10" \
        -e "BUBBLE_NUM_THREADS=8" \
        "$FULL_IMAGE_NAME"
    
    echo ""
    echo "All tests completed!"
}

# Function to clean up
clean_image() {
    echo "Removing Docker image: $FULL_IMAGE_NAME"
    docker rmi "$FULL_IMAGE_NAME"
}

# Main script logic
case "${1:-}" in
    build)
        build_image
        ;;
    run)
        shift
        run_container "$@"
        ;;
    test)
        run_tests
        ;;
    clean)
        clean_image
        ;;
    ""|-h|--help)
        show_usage
        ;;
    *)
        echo "Unknown command: $1"
        show_usage
        exit 1
        ;;
esac