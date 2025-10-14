# Bubble Memory Stress Program Makefile

CC = gcc
CFLAGS = -O3 -march=native -mtune=native -Wall -Wextra -std=c99 -fopenmp
LDFLAGS = -fopenmp
TARGET = bubble
SOURCE = bubble.c

# Default target
all: $(TARGET)

# Build the bubble program
$(TARGET): $(SOURCE)
	$(CC) $(CFLAGS) -o $(TARGET) $(SOURCE) $(LDFLAGS)

# Build with debug symbols
debug: CFLAGS += -g -DDEBUG -O0
debug: $(TARGET)

# Build optimized version for release
release: CFLAGS += -DNDEBUG -flto
release: $(TARGET)

# Clean build artifacts
clean:
	rm -f $(TARGET)

# Install to system (requires sudo)
install: $(TARGET)
	install -m 755 $(TARGET) /usr/local/bin/
	install -m 755 run_bubble.sh /usr/local/bin/
	install -m 644 bubble.conf /etc/

# Uninstall from system (requires sudo)
uninstall:
	rm -f /usr/local/bin/$(TARGET)
	rm -f /usr/local/bin/run_bubble.sh
	rm -f /etc/bubble.conf

# Test build
test: $(TARGET)
	./$(TARGET) 0.1

# Run with low pressure
test-low: $(TARGET)
	BUBBLE_PRESSURE_LEVEL=0.1 BUBBLE_RUNTIME_SECONDS=5 ./run_bubble.sh

# Run with medium pressure
test-medium: $(TARGET)
	BUBBLE_PRESSURE_LEVEL=0.5 BUBBLE_RUNTIME_SECONDS=5 ./run_bubble.sh

# Run with high pressure
test-high: $(TARGET)
	BUBBLE_PRESSURE_LEVEL=0.9 BUBBLE_RUNTIME_SECONDS=5 ./run_bubble.sh

# Docker build
docker-build:
	docker build -t bubble .

# Docker run example
docker-run: docker-build
	docker run --rm -e BUBBLE_PRESSURE_LEVEL=0.5 bubble

.PHONY: all debug release clean install uninstall test test-low test-medium test-high docker-build docker-run