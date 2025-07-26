# Proof of Concept / Testbed Validation for c6620 Node
## Setup
```bash
sudo bash ./script/types/c6620/001-install-rdt.sh 
sudo bash ./script/types/c6620/002-install-runtime.sh
sudo bash ./script/types/c6620/003-setup-rdt.sh
```

## Spin up limited and reference container
This container will be assigned to a custom CLOS group to limit its L3 cache.
```bash
docker run -d --name ltd     --cpuset-cpus 0-31         --cap-add SYS_ADMIN --cap-add PERFMON   ubuntu:22.04 sleep infinity
```
We save the container ID of the to be limited container to later pass it to `runc`:
```bash
CID=$(docker inspect --format '{{.Id}}' ltd)
```
and start a second container for reference which will be unchanged in terms of resources
```bash
docker run -d --name ref     --cpuset-cpus 0-31         --cap-add SYS_ADMIN --cap-add PERFMON   ubuntu:22.04 sleep infinity
```
We install `stress-ng` as a benchmark and `perf` to access the L3 miss rate later on.

```bash
for C in ltd ref; do   docker exec -u root $C bash -c     "apt-get update -qq && \
     DEBIAN_FRONTEND=noninteractive apt-get install -y -qq stress-ng linux-tools-common linux-tools-generic linux-tools-`uname -r`"; done
```	 
We assign the `ltd` container to a custom CLOS class which will use the bitmask `7`:
```bash
sudo runc      --root /run/docker/runtime-runc/moby      update --l3-cache-schema "L3:0=7" "$CID"
```
after, we can check the schemata:
```bash
cat /sys/fs/resctrl/$CID/schemata
```
and start the benchmarks

```bash
docker exec ltd stress-ng --cache 4 --cache-level 3      --timeout 30s --metrics-brief --perf
```
and
```bash
docker exec ref stress-ng --cache 4 --cache-level 3      --timeout 30s --metrics-brief --perf
```