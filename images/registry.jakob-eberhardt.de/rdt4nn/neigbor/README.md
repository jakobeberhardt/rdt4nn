# Neighbor Workload Container

This container provides a configurable memory workload generator that can be used as a "noisy neighbor" in container benchmarking scenarios.

## Features

- **Configurable Duration**: Set how long the workload runs
- **Configurable Buffer Size**: Control memory usage (supports KB, MB, GB)
- **Multiple Access Patterns**: 
  - `random`: Random memory access
  - `stride`: Strided memory access (4KB stride)
  - `sequential`: Sequential memory access

```shell
/neighbor --duration 3600 --buffer-size 1000MB --pattern random" container=random

# Huge pages (best-effort THP)
/neighbor --duration 3600 --buffer-size 1000MB --pattern random --hugepages thp

# Explicit 1GB hugetlb hugepages (requires host reservation; mapping is rounded up to 1GB)
/neighbor --duration 3600 --buffer-size 1GB --pattern random --hugepages hugetlb-1g

# Small logical buffer, but backed by a 1GB hugetlb page (still consumes one full 1GB hugepage once touched)
/neighbor --duration 3600 --buffer-size 6MB --pattern random --hugepages hugetlb-1g
```
