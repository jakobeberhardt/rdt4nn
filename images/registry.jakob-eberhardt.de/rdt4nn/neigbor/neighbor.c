#define _GNU_SOURCE

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <time.h>
#include <signal.h>
#include <errno.h>
#include <stdint.h>
#include <sys/mman.h>
#ifdef __linux__
#include <linux/mman.h>
#endif

static volatile int running = 1;

typedef enum {
    ALLOC_MALLOC = 0,
    ALLOC_THP,
    ALLOC_HUGETLB_2M,
    ALLOC_HUGETLB_1G,
} alloc_mode_t;

typedef struct {
    char *ptr;
    long size;
    alloc_mode_t mode;
} buffer_alloc_t;

void signal_handler(int sig) {
    running = 0;
}

void print_usage(const char *prog_name) {
    printf("Usage: %s [OPTIONS]\n", prog_name);
    printf("Options:\n");
    printf("  --duration SECONDS     Run for specified duration (default: 3600)\n");
    printf("  --buffer-size SIZE     Buffer size (e.g., 500MB, 1GB) (default: 100MB)\n");
    printf("  --pattern PATTERN      Access pattern: random/stride/sequential (default: random)\n");
    printf("  --hugepages MODE       Memory backing: none/thp/hugetlb-2m/hugetlb-1g (default: none)\n");
    printf("  --help                 Show this help message\n");
}

#ifndef MAP_HUGETLB
#define MAP_HUGETLB 0x40000
#endif

#ifndef MAP_HUGE_SHIFT
#define MAP_HUGE_SHIFT 26
#endif

#ifndef MAP_HUGE_2MB
#define MAP_HUGE_2MB (21 << MAP_HUGE_SHIFT)
#endif

#ifndef MAP_HUGE_1GB
#define MAP_HUGE_1GB (30 << MAP_HUGE_SHIFT)
#endif

static const char *alloc_mode_to_string(alloc_mode_t mode) {
    switch (mode) {
        case ALLOC_MALLOC: return "none";
        case ALLOC_THP: return "thp";
        case ALLOC_HUGETLB_2M: return "hugetlb-2m";
        case ALLOC_HUGETLB_1G: return "hugetlb-1g";
        default: return "unknown";
    }
}

static int parse_alloc_mode(const char *mode_str, alloc_mode_t *out_mode) {
    if (mode_str == NULL || out_mode == NULL) {
        return -1;
    }
    if (strcmp(mode_str, "none") == 0) {
        *out_mode = ALLOC_MALLOC;
        return 0;
    }
    if (strcmp(mode_str, "thp") == 0) {
        *out_mode = ALLOC_THP;
        return 0;
    }
    if (strcmp(mode_str, "hugetlb-2m") == 0) {
        *out_mode = ALLOC_HUGETLB_2M;
        return 0;
    }
    if (strcmp(mode_str, "hugetlb-1g") == 0) {
        *out_mode = ALLOC_HUGETLB_1G;
        return 0;
    }
    return -1;
}

static int alloc_buffer(long buffer_size, alloc_mode_t mode, buffer_alloc_t *out_alloc) {
    if (out_alloc == NULL || buffer_size <= 0) {
        errno = EINVAL;
        return -1;
    }

    out_alloc->ptr = NULL;
    out_alloc->size = buffer_size;
    out_alloc->mode = mode;

    if (mode == ALLOC_MALLOC) {
        out_alloc->ptr = (char *)malloc((size_t)buffer_size);
        return out_alloc->ptr ? 0 : -1;
    }

    int flags = MAP_PRIVATE | MAP_ANONYMOUS;
    if (mode == ALLOC_HUGETLB_2M || mode == ALLOC_HUGETLB_1G) {
        flags |= MAP_HUGETLB;
        if (mode == ALLOC_HUGETLB_2M) {
            flags |= MAP_HUGE_2MB;
            if ((buffer_size % (2L * 1024 * 1024)) != 0) {
                errno = EINVAL;
                return -1;
            }
        } else {
            flags |= MAP_HUGE_1GB;
            if ((buffer_size % (1024L * 1024 * 1024)) != 0) {
                errno = EINVAL;
                return -1;
            }
        }
    }

    void *p = mmap(NULL, (size_t)buffer_size, PROT_READ | PROT_WRITE, flags, -1, 0);
    if (p == MAP_FAILED) {
        out_alloc->ptr = NULL;
        return -1;
    }
    out_alloc->ptr = (char *)p;

    if (mode == ALLOC_THP) {
        // Best-effort: encourages transparent huge pages (usually 2MB; 1GB THP depends on kernel settings).
#ifdef MADV_HUGEPAGE
        (void)madvise(out_alloc->ptr, (size_t)buffer_size, MADV_HUGEPAGE);
#endif
    }
    return 0;
}

static void free_buffer(buffer_alloc_t *alloc) {
    if (alloc == NULL || alloc->ptr == NULL) {
        return;
    }
    if (alloc->mode == ALLOC_MALLOC) {
        free(alloc->ptr);
    } else {
        (void)munmap(alloc->ptr, (size_t)alloc->size);
    }
    alloc->ptr = NULL;
}

long parse_size(const char *size_str) {
    char *endptr;
    double size = strtod(size_str, &endptr);
    
    if (endptr == size_str) {
        fprintf(stderr, "Invalid size format: %s\n", size_str);
        return -1;
    }
    
    long multiplier = 1;
    if (strlen(endptr) > 0) {
        if (strcasecmp(endptr, "KB") == 0) {
            multiplier = 1024;
        } else if (strcasecmp(endptr, "MB") == 0) {
            multiplier = 1024 * 1024;
        } else if (strcasecmp(endptr, "GB") == 0) {
            multiplier = 1024 * 1024 * 1024;
        } else {
            fprintf(stderr, "Unknown size suffix: %s\n", endptr);
            return -1;
        }
    }
    
    return (long)(size * multiplier);
}

void random_access(char *buffer, long size, int duration) {
    time_t start_time = time(NULL);
    srand(time(NULL));
    long operations = 0;
    
    printf("Starting random access pattern for %d seconds with buffer size %ld bytes\n", duration, size);
    
    while (running && (time(NULL) - start_time) < duration) {
        // Access multiple random locations per iteration for intensive memory pressure
        for (int burst = 0; burst < 10000 && running; burst++) {
            long offset = rand() % size;
            volatile char temp = buffer[offset];
            buffer[offset] = (char)((temp + offset) % 256);
            operations++;
        }
    }
    
    printf("Random access completed: %ld operations\n", operations);
}

void stride_access(char *buffer, long size, int duration) {
    time_t start_time = time(NULL);
    long stride = 4096; // 4KB stride for cache-unfriendly access
    long pos = 0;
    long operations = 0;
    
    printf("Starting stride access pattern for %d seconds with buffer size %ld bytes\n", duration, size);
    
    while (running && (time(NULL) - start_time) < duration) {
        // Access multiple stride locations per iteration
        for (int burst = 0; burst < 1000 && running; burst++) {
            volatile char temp = buffer[pos];
            buffer[pos] = (char)((temp + pos) % 256);
            pos = (pos + stride) % size;
            operations++;
        }
    }
    
    printf("Stride access completed: %ld operations\n", operations);
}

void sequential_access(char *buffer, long size, int duration) {
    time_t start_time = time(NULL);
    long pos = 0;
    long operations = 0;
    long total_passes = 0;
    
    printf("Starting sequential access pattern for %d seconds with buffer size %ld bytes\n", duration, size);
    
    while (running && (time(NULL) - start_time) < duration) {
        for (long i = 0; i < size && running && (time(NULL) - start_time) < duration; i += sizeof(long)) {
            volatile long *ptr = (long*)(buffer + i);
            *ptr = (*ptr + i) % 0xFFFFFFFF;
            operations++;
        }
        total_passes++;
        printf("Sequential pass %ld completed\n", total_passes);
    }
    
    printf("Sequential access completed: %ld operations, %ld passes\n", operations, total_passes);
}

int main(int argc, char *argv[]) {
    int duration = 3600; 
    long buffer_size = 100 * 1024 * 1024; 
    char pattern[32] = "random"; 
    alloc_mode_t alloc_mode = ALLOC_MALLOC;
    
    signal(SIGTERM, signal_handler);
    signal(SIGINT, signal_handler);
    
    for (int i = 1; i < argc; i++) {
        if (strcmp(argv[i], "--duration") == 0) {
            if (i + 1 < argc) {
                duration = atoi(argv[++i]);
                if (duration <= 0) {
                    fprintf(stderr, "Invalid duration: %s\n", argv[i]);
                    return 1;
                }
            } else {
                fprintf(stderr, "--duration requires a value\n");
                return 1;
            }
        } else if (strcmp(argv[i], "--buffer-size") == 0) {
            if (i + 1 < argc) {
                buffer_size = parse_size(argv[++i]);
                if (buffer_size <= 0) {
                    fprintf(stderr, "Invalid buffer size: %s\n", argv[i]);
                    return 1;
                }
            } else {
                fprintf(stderr, "--buffer-size requires a value\n");
                return 1;
            }
        } else if (strcmp(argv[i], "--pattern") == 0) {
            if (i + 1 < argc) {
                strncpy(pattern, argv[++i], sizeof(pattern) - 1);
                pattern[sizeof(pattern) - 1] = '\0';
                if (strcmp(pattern, "random") != 0 && 
                    strcmp(pattern, "stride") != 0 && 
                    strcmp(pattern, "sequential") != 0) {
                    fprintf(stderr, "Invalid pattern: %s. Use random, stride, or sequential\n", pattern);
                    return 1;
                }
            } else {
                fprintf(stderr, "--pattern requires a value\n");
                return 1;
            }
        } else if (strcmp(argv[i], "--help") == 0) {
            print_usage(argv[0]);
            return 0;
        } else if (strcmp(argv[i], "--hugepages") == 0) {
            if (i + 1 < argc) {
                if (parse_alloc_mode(argv[++i], &alloc_mode) != 0) {
                    fprintf(stderr, "Invalid --hugepages mode: %s. Use none, thp, hugetlb-2m, or hugetlb-1g\n", argv[i]);
                    return 1;
                }
            } else {
                fprintf(stderr, "--hugepages requires a value\n");
                return 1;
            }
        } else {
            fprintf(stderr, "Unknown option: %s\n", argv[i]);
            print_usage(argv[0]);
            return 1;
        }
    }
    
    printf("Neighbor workload starting...\n");
    printf("Duration: %d seconds\n", duration);
    printf("Buffer size: %ld bytes (%.2f MB)\n", buffer_size, (double)buffer_size / (1024 * 1024));
    printf("Access pattern: %s\n", pattern);
    printf("Hugepages: %s\n", alloc_mode_to_string(alloc_mode));

    if (alloc_mode == ALLOC_HUGETLB_2M && (buffer_size % (2L * 1024 * 1024)) != 0) {
        fprintf(stderr, "Buffer size must be a multiple of 2MB for --hugepages hugetlb-2m\n");
        return 1;
    }
    if (alloc_mode == ALLOC_HUGETLB_1G && (buffer_size % (1024L * 1024 * 1024)) != 0) {
        fprintf(stderr, "Buffer size must be a multiple of 1GB for --hugepages hugetlb-1g\n");
        return 1;
    }

    buffer_alloc_t alloc;
    if (alloc_buffer(buffer_size, alloc_mode, &alloc) != 0) {
        fprintf(stderr, "Failed to allocate %ld bytes (hugepages=%s): %s\n", buffer_size, alloc_mode_to_string(alloc_mode), strerror(errno));
        return 1;
    }
    char *buffer = alloc.ptr;
    
    printf("Initializing buffer and forcing allocation...\n");
    for (long i = 0; i < buffer_size; i += 4096) {
        buffer[i] = (char)(i % 256);
    }
    buffer[buffer_size - 1] = 0xFF;
    printf("Buffer initialization completed\n");
    
    if (strcmp(pattern, "random") == 0) {
        random_access(buffer, buffer_size, duration);
    } else if (strcmp(pattern, "stride") == 0) {
        stride_access(buffer, buffer_size, duration);
    } else if (strcmp(pattern, "sequential") == 0) {
        sequential_access(buffer, buffer_size, duration);
    }
    
    printf("Neighbor workload completed\n");
    free_buffer(&alloc);
    return 0;
}
