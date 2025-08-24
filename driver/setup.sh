#!/bin/bash

set -e

echo "RDT4NN Driver Setup Script"
echo "========================="

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if running as root for some operations
check_sudo() {
    if [ "$EUID" -eq 0 ]; then
        print_warning "Running as root. Some operations will be skipped."
        return 0
    else
        return 1
    fi
}

# Install Go if not present
install_go() {
    if command -v go &> /dev/null; then
        print_status "Go is already installed: $(go version)"
        return 0
    fi

    print_status "Installing Go..."
    
    # Download and install Go
    GO_VERSION="1.21.5"
    GO_OS="linux"
    GO_ARCH="amd64"
    
    if [ "$(uname -m)" = "aarch64" ]; then
        GO_ARCH="arm64"
    fi
    
    GO_TARBALL="go${GO_VERSION}.${GO_OS}-${GO_ARCH}.tar.gz"
    
    cd /tmp
    wget -q "https://golang.org/dl/${GO_TARBALL}"
    
    if check_sudo; then
        tar -C /usr/local -xzf "$GO_TARBALL"
    else
        sudo tar -C /usr/local -xzf "$GO_TARBALL"
    fi
    
    # Add Go to PATH
    if ! grep -q "/usr/local/go/bin" ~/.bashrc; then
        echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
        echo 'export GOPATH=$HOME/go' >> ~/.bashrc
        echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.bashrc
    fi
    
    # For current session
    export PATH=$PATH:/usr/local/go/bin
    export GOPATH=$HOME/go
    export PATH=$PATH:$GOPATH/bin
    
    print_status "Go installed successfully"
    
    rm -f "/tmp/$GO_TARBALL"
}

# Install system dependencies
install_system_deps() {
    print_status "Installing system dependencies..."
    
    # Update package list
    if check_sudo; then
        apt update
    else
        sudo apt update
    fi
    
    # Install required packages
    PACKAGES="linux-tools-common linux-tools-generic linux-tools-`uname -r` curl wget build-essential"
    
    # Optional: Intel RDT tools (may not be available on all systems)
    # PACKAGES="$PACKAGES intel-pqos-tools"
    
    if check_sudo; then
        apt install -y $PACKAGES
    else
        sudo apt install -y $PACKAGES
    fi
    
    # Add user to docker group
    if ! check_sudo && ! groups | grep -q docker; then
        sudo usermod -aG docker "$USER"
        print_warning "Added $USER to docker group. Please logout and login again for this to take effect."
    fi
}

# Setup Intel RDT
setup_rdt() {
    print_status "Setting up Intel RDT..."
    
    # Check if RDT is supported
    if ! grep -q "rdt" /proc/cpuinfo; then
        print_warning "Intel RDT may not be supported on this CPU"
        return 0
    fi
    
    # Create and mount resctrl filesystem
    if check_sudo; then
        mkdir -p /sys/fs/resctrl
        mount -t resctrl resctrl /sys/fs/resctrl 2>/dev/null || print_warning "resctrl already mounted"
        chmod -R 755 /sys/fs/resctrl 2>/dev/null || true
    else
        sudo mkdir -p /sys/fs/resctrl
        sudo mount -t resctrl resctrl /sys/fs/resctrl 2>/dev/null || print_warning "resctrl already mounted"
        sudo chmod -R 755 /sys/fs/resctrl 2>/dev/null || true
    fi
    
    print_status "RDT setup complete"
}

# Setup perf events
setup_perf() {
    print_status "Setting up perf events..."
    
    # Allow perf events for non-root users
    if check_sudo; then
        echo -1 > /proc/sys/kernel/perf_event_paranoid
    else
        echo -1 | sudo tee /proc/sys/kernel/perf_event_paranoid > /dev/null
    fi
    
    print_status "Perf events setup complete"
}

# Install Docker if not present
install_docker() {
    if command -v docker &> /dev/null; then
        print_status "Docker is already installed: $(docker --version)"
        return 0
    fi
    
    print_status "Installing Docker..."
    
    # Install Docker using convenience script
    curl -fsSL https://get.docker.com -o get-docker.sh
    if check_sudo; then
        sh get-docker.sh
    else
        sudo sh get-docker.sh
    fi
    
    rm -f get-docker.sh
    
    # Start Docker service
    if check_sudo; then
        systemctl start docker
        systemctl enable docker
    else
        sudo systemctl start docker
        sudo systemctl enable docker
    fi
    
    print_status "Docker installed successfully"
}

# Setup InfluxDB (optional)
setup_influxdb() {
    if [ "$1" = "--skip-influxdb" ]; then
        print_status "Skipping InfluxDB setup"
        return 0
    fi
    
    print_status "Setting up InfluxDB..."
    
    # Install InfluxDB using Docker
    docker run -d \
        --name influxdb \
        -p 8086:8086 \
        -e DOCKER_INFLUXDB_INIT_MODE=setup \
        -e DOCKER_INFLUXDB_INIT_USERNAME=jakob \
        -e DOCKER_INFLUXDB_INIT_PASSWORD=123 \
        -e DOCKER_INFLUXDB_INIT_ORG=rdt4nn \
        -e DOCKER_INFLUXDB_INIT_BUCKET=benchmarks \
        -v influxdb2-data:/var/lib/influxdb2 \
        -v influxdb2-config:/etc/influxdb2 \
        influxdb:2.7-alpine 2>/dev/null || print_warning "InfluxDB container already exists or failed to start"
    
    print_status "InfluxDB setup complete (running on port 8086)"
}

# Build the RDT4NN driver
build_driver() {
    print_status "Building RDT4NN driver..."
    
    # Initialize Go module if not already done
    if [ ! -f go.mod ]; then
        go mod init github.com/jakobeberhardt/rdt4nn/driver
    fi
    
    # Download dependencies
    go mod tidy
    
    # Build the driver
    go build -o rdt4nn-driver .
    
    print_status "Driver built successfully: ./rdt4nn-driver"
}

# Verify installation
verify_installation() {
    print_status "Verifying installation..."
    
    # Check Go
    if ! command -v go &> /dev/null; then
        print_error "Go installation failed"
        return 1
    fi
    
    # Check Docker
    if ! command -v docker &> /dev/null; then
        print_error "Docker installation failed"
        return 1
    fi
    
    # Check if driver was built
    if [ ! -f "./rdt4nn-driver" ]; then
        print_error "Driver build failed"
        return 1
    fi
    
    # Test driver
    ./rdt4nn-driver version || print_warning "Driver version check failed"
    
    # Validate example configs
    ./rdt4nn-driver validate -c examples/simple_test.yml || print_warning "Example validation failed"
    
    print_status "Installation verification complete"
}

# Main setup function
main() {
    echo "Starting setup process..."
    
    # Parse command line arguments
    SKIP_INFLUXDB=false
    SKIP_BUILD=false
    
    for arg in "$@"; do
        case $arg in
            --skip-influxdb)
                SKIP_INFLUXDB=true
                ;;
            --skip-build)
                SKIP_BUILD=true
                ;;
            --help|-h)
                echo "Usage: $0 [--skip-influxdb] [--skip-build] [--help]"
                echo "  --skip-influxdb  Skip InfluxDB setup"
                echo "  --skip-build     Skip building the driver"
                echo "  --help           Show this help message"
                exit 0
                ;;
        esac
    done
    
    # Install system dependencies
    install_system_deps
    
    # Install Go
    install_go
    
    # Install Docker
    install_docker
    
    # Setup RDT
    #setup_rdt
    
    # Setup perf
    setup_perf
    
    # Setup InfluxDB
    if [ "$SKIP_INFLUXDB" = false ]; then
        setup_influxdb
    else
        setup_influxdb --skip-influxdb
    fi
    
    # Build driver
    if [ "$SKIP_BUILD" = false ]; then
        build_driver
    fi
    
    # Verify installation
    verify_installation
    
    print_status "Setup complete!"
    echo ""
    echo "Next steps:"
    echo "1. Logout and login again to activate Docker group membership"
    echo "2. Source your bashrc: source ~/.bashrc"
    echo "3. Test the driver: ./rdt4nn-driver validate -c examples/simple_test.yml"
    echo "4. Run a benchmark: ./rdt4nn-driver -c examples/simple_test.yml"
}

# Run main function
main "$@"
