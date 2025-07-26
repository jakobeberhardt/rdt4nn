#!/bin/bash
set -euo pipefail

if [[ $EUID -ne 0 ]]; then
  echo "Please run as root (sudo)." >&2
  exit 1
fi

add-apt-repository -y universe    

apt-get update
apt-get install -y \
    intel-cmt-cat   \
    msr-tools        

modprobe msr                                    
echo msr > /etc/modules-load.d/intel_rdt.conf  
mkdir -p /sys/fs/resctrl                        
if ! mountpoint -q /sys/fs/resctrl; then
  mount -t resctrl resctrl /sys/fs/resctrl
fi
grep -q '^resctrl' /etc/fstab || \
  echo 'resctrl /sys/fs/resctrl resctrl defaults 0 0' >> /etc/fstab

echo
echo "== Installed versions =="
pqos -V        
rdtset -V || true
echo
echo "Intel RDT tooling ready. Try 'pqos -s' to view caps."