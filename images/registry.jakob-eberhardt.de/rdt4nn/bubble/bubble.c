#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <omp.h>
#include <time.h>
#include <getopt.h>
#include <signal.h>

#define LFSR_MASK 0xd0000001u
#define MB_TO_BYTES(mb) ((mb) * 1024 * 1024)
#define DUMP_SIZE 100
#define CHECK_INTERVAL 10000

static volatile sig_atomic_t keep_running = 1;

void signal_handler(int signum) {
    keep_running = 0;
}

typedef struct {
    size_t min_working_size_mb;
    size_t max_working_size_mb;
    int ramp_time_sec;
    int num_threads;
    int streaming_ratio;
} bubble_config_t;

static inline unsigned lfsr_rand(unsigned *lfsr) {
    *lfsr = (*lfsr >> 1) ^ (unsigned)((0 - (*lfsr & 1u)) & LFSR_MASK);
    return *lfsr;
}

void random_access_kernel(int *data_chunk, size_t footprint_size, unsigned *lfsr) {
    int dump[DUMP_SIZE] = {0};
    unsigned long iter = 0;
    
    while (1) {
        unsigned r = lfsr_rand(lfsr) % footprint_size;
        
        dump[0] += data_chunk[r]++;
        dump[1] += data_chunk[r]++;
        dump[2] += data_chunk[r]++;
        dump[3] += data_chunk[r]++;
        dump[4] += data_chunk[r]++;
        dump[5] += data_chunk[r]++;
        dump[6] += data_chunk[r]++;
        dump[7] += data_chunk[r]++;
        dump[8] += data_chunk[r]++;
        dump[9] += data_chunk[r]++;
        dump[10] += data_chunk[r]++;
        dump[11] += data_chunk[r]++;
        dump[12] += data_chunk[r]++;
        dump[13] += data_chunk[r]++;
        dump[14] += data_chunk[r]++;
        dump[15] += data_chunk[r]++;
        dump[16] += data_chunk[r]++;
        dump[17] += data_chunk[r]++;
        dump[18] += data_chunk[r]++;
        dump[19] += data_chunk[r]++;
        dump[20] += data_chunk[r]++;
        dump[21] += data_chunk[r]++;
        dump[22] += data_chunk[r]++;
        dump[23] += data_chunk[r]++;
        dump[24] += data_chunk[r]++;
        dump[25] += data_chunk[r]++;
        dump[26] += data_chunk[r]++;
        dump[27] += data_chunk[r]++;
        dump[28] += data_chunk[r]++;
        dump[29] += data_chunk[r]++;
        dump[30] += data_chunk[r]++;
        dump[31] += data_chunk[r]++;
        dump[32] += data_chunk[r]++;
        dump[33] += data_chunk[r]++;
        dump[34] += data_chunk[r]++;
        dump[35] += data_chunk[r]++;
        dump[36] += data_chunk[r]++;
        dump[37] += data_chunk[r]++;
        dump[38] += data_chunk[r]++;
        dump[39] += data_chunk[r]++;
        dump[40] += data_chunk[r]++;
        dump[41] += data_chunk[r]++;
        dump[42] += data_chunk[r]++;
        dump[43] += data_chunk[r]++;
        dump[44] += data_chunk[r]++;
        dump[45] += data_chunk[r]++;
        dump[46] += data_chunk[r]++;
        dump[47] += data_chunk[r]++;
        dump[48] += data_chunk[r]++;
        dump[49] += data_chunk[r]++;
        dump[50] += data_chunk[r]++;
        dump[51] += data_chunk[r]++;
        dump[52] += data_chunk[r]++;
        dump[53] += data_chunk[r]++;
        dump[54] += data_chunk[r]++;
        dump[55] += data_chunk[r]++;
        dump[56] += data_chunk[r]++;
        dump[57] += data_chunk[r]++;
        dump[58] += data_chunk[r]++;
        dump[59] += data_chunk[r]++;
        dump[60] += data_chunk[r]++;
        dump[61] += data_chunk[r]++;
        dump[62] += data_chunk[r]++;
        dump[63] += data_chunk[r]++;
        dump[64] += data_chunk[r]++;
        dump[65] += data_chunk[r]++;
        dump[66] += data_chunk[r]++;
        dump[67] += data_chunk[r]++;
        dump[68] += data_chunk[r]++;
        dump[69] += data_chunk[r]++;
        dump[70] += data_chunk[r]++;
        dump[71] += data_chunk[r]++;
        dump[72] += data_chunk[r]++;
        dump[73] += data_chunk[r]++;
        dump[74] += data_chunk[r]++;
        dump[75] += data_chunk[r]++;
        dump[76] += data_chunk[r]++;
        dump[77] += data_chunk[r]++;
        dump[78] += data_chunk[r]++;
        dump[79] += data_chunk[r]++;
        dump[80] += data_chunk[r]++;
        dump[81] += data_chunk[r]++;
        dump[82] += data_chunk[r]++;
        dump[83] += data_chunk[r]++;
        dump[84] += data_chunk[r]++;
        dump[85] += data_chunk[r]++;
        dump[86] += data_chunk[r]++;
        dump[87] += data_chunk[r]++;
        dump[88] += data_chunk[r]++;
        dump[89] += data_chunk[r]++;
        dump[90] += data_chunk[r]++;
        dump[91] += data_chunk[r]++;
        dump[92] += data_chunk[r]++;
        dump[93] += data_chunk[r]++;
        dump[94] += data_chunk[r]++;
        dump[95] += data_chunk[r]++;
        dump[96] += data_chunk[r]++;
        dump[97] += data_chunk[r]++;
        dump[98] += data_chunk[r]++;
        dump[99] += data_chunk[r]++;
        
        if (++iter % CHECK_INTERVAL == 0 && !keep_running) {
            break;
        }
    }
}

void streaming_kernel(double *bw_data, size_t stream_size) {
    double scalar = 3.0;
    unsigned long iter = 0;
    
    while (1) {
        double *mid = bw_data + (stream_size / 2);
        
        for (size_t i = 0; i < stream_size / 2; i++) {
            bw_data[i] = scalar * mid[i];
        }
        
        for (size_t i = 0; i < stream_size / 2; i++) {
            mid[i] = scalar * bw_data[i];
        }
        
        if (++iter % CHECK_INTERVAL == 0 && !keep_running) {
            break;
        }
    }
}

void run_bubble(bubble_config_t *config) {
    size_t current_size_mb = config->min_working_size_mb;
    size_t max_size_mb = config->max_working_size_mb;
    int ramp_steps = config->ramp_time_sec > 0 ? config->ramp_time_sec : 0;
    size_t step_size_mb = ramp_steps > 0 ? (max_size_mb - current_size_mb) / ramp_steps : 0;
    
    signal(SIGINT, signal_handler);
    signal(SIGTERM, signal_handler);
    
    omp_set_num_threads(config->num_threads);
    
    printf("Bubble Configuration:\n");
    printf("  Threads: %d\n", config->num_threads);
    printf("  Working Set: %zu MB -> %zu MB\n", config->min_working_size_mb, config->max_working_size_mb);
    if (config->ramp_time_sec > 0) {
        printf("  Ramp Time: %d seconds (will exit after)\n", config->ramp_time_sec);
    } else {
        printf("  Ramp Time: none (runs indefinitely)\n");
    }
    printf("  Streaming Ratio: %d%%\n", config->streaming_ratio);
    printf("\n");
    fflush(stdout);
    
    if (ramp_steps > 0) {
        for (int step = 0; step <= ramp_steps && keep_running; step++) {
            if (step > 0) {
                current_size_mb += step_size_mb;
                if (current_size_mb > max_size_mb) {
                    current_size_mb = max_size_mb;
                }
            }
            
            size_t current_bytes = MB_TO_BYTES(current_size_mb);
            size_t int_elements = current_bytes / sizeof(int);
            size_t double_elements = current_bytes / sizeof(double);
            
            printf("Step %d/%d: Working set = %zu MB\n", step, ramp_steps, current_size_mb);
            fflush(stdout);
            
            int *random_data = (int *)malloc(current_bytes);
            double *stream_data = (double *)malloc(current_bytes);
            
            if (!random_data || !stream_data) {
                fprintf(stderr, "Memory allocation failed for %zu MB\n", current_size_mb);
                exit(1);
            }
            
            memset(random_data, 0, current_bytes);
            memset(stream_data, 0, current_bytes);
            
            time_t start_time = time(NULL);
            
            #pragma omp parallel
            {
                int tid = omp_get_thread_num();
                unsigned lfsr = 0xACE1u + tid;
                
                int do_streaming = (tid * 100 / config->num_threads) < config->streaming_ratio;
                
                while (keep_running && (time(NULL) - start_time < 1)) {
                    if (do_streaming) {
                        double *mid = stream_data + (double_elements / 2);
                        double scalar = 3.0;
                        for (size_t i = 0; i < double_elements / 2; i++) {
                            stream_data[i] = scalar * mid[i];
                        }
                        for (size_t i = 0; i < double_elements / 2; i++) {
                            mid[i] = scalar * stream_data[i];
                        }
                    } else {
                        int dump[DUMP_SIZE] = {0};
                        for (int j = 0; j < 1000; j++) {
                            unsigned r = lfsr_rand(&lfsr) % int_elements;
                            for (int k = 0; k < DUMP_SIZE; k++) {
                                dump[k] += random_data[r]++;
                            }
                        }
                    }
                }
            }
            
            free(random_data);
            free(stream_data);
        }
        printf("\nRamp complete, exiting.\n");
    } else {
        current_size_mb = max_size_mb;
        size_t current_bytes = MB_TO_BYTES(current_size_mb);
        size_t int_elements = current_bytes / sizeof(int);
        size_t double_elements = current_bytes / sizeof(double);
        
        printf("Running at %zu MB (no ramp)\n", current_size_mb);
        fflush(stdout);
        
        int *random_data = (int *)malloc(current_bytes);
        double *stream_data = (double *)malloc(current_bytes);
        
        if (!random_data || !stream_data) {
            fprintf(stderr, "Memory allocation failed for %zu MB\n", current_size_mb);
            exit(1);
        }
        
        memset(random_data, 0, current_bytes);
        memset(stream_data, 0, current_bytes);
        
        #pragma omp parallel
        {
            int tid = omp_get_thread_num();
            unsigned lfsr = 0xACE1u + tid;
            
            int do_streaming = (tid * 100 / config->num_threads) < config->streaming_ratio;
            
            if (do_streaming) {
                streaming_kernel(stream_data, double_elements);
            } else {
                random_access_kernel(random_data, int_elements, &lfsr);
            }
        }
        
        free(random_data);
        free(stream_data);
    }
    
    if (!keep_running) {
        printf("\nBubble terminated by signal\n");
    }
}

void print_usage(const char *prog_name) {
    printf("Usage: %s [OPTIONS]\n", prog_name);
    printf("\nMemory Pressure Bubble for Container Interference Studies\n\n");
    printf("Options:\n");
    printf("  --min-size MB          Minimum working set size in MB (default: 1)\n");
    printf("  --max-size MB          Maximum working set size in MB (default: 100)\n");
    printf("  --ramp-time SECONDS    Time to ramp from min to max in seconds (default: 0)\n");
    printf("  --threads N            Number of threads (default: system cores)\n");
    printf("  --streaming-ratio %%    Percentage of threads doing streaming (default: 50)\n");
    printf("  -h, --help             Show this help message\n");
    printf("\nExamples:\n");
    printf("  %s --max-size 1024\n", prog_name);
    printf("  %s --min-size 10 --max-size 1000 --ramp-time 60\n", prog_name);
    printf("  %s --threads 4 --streaming-ratio 25\n", prog_name);
}

int main(int argc, char *argv[]) {
    bubble_config_t config = {
        .min_working_size_mb = 1,
        .max_working_size_mb = 100,
        .ramp_time_sec = 0,
        .num_threads = omp_get_max_threads(),
        .streaming_ratio = 50
    };
    
    static struct option long_options[] = {
        {"min-size", required_argument, 0, 'm'},
        {"max-size", required_argument, 0, 'M'},
        {"ramp-time", required_argument, 0, 'r'},
        {"threads", required_argument, 0, 't'},
        {"streaming-ratio", required_argument, 0, 's'},
        {"help", no_argument, 0, 'h'},
        {0, 0, 0, 0}
    };
    
    int opt;
    int option_index = 0;
    
    while ((opt = getopt_long(argc, argv, "h", long_options, &option_index)) != -1) {
        switch (opt) {
            case 'm':
                config.min_working_size_mb = atoi(optarg);
                break;
            case 'M':
                config.max_working_size_mb = atoi(optarg);
                break;
            case 'r':
                config.ramp_time_sec = atoi(optarg);
                break;
            case 't':
                config.num_threads = atoi(optarg);
                break;
            case 's':
                config.streaming_ratio = atoi(optarg);
                if (config.streaming_ratio < 0) config.streaming_ratio = 0;
                if (config.streaming_ratio > 100) config.streaming_ratio = 100;
                break;
            case 'h':
                print_usage(argv[0]);
                return 0;
            default:
                print_usage(argv[0]);
                return 1;
        }
    }
    
    if (config.min_working_size_mb > config.max_working_size_mb) {
        fprintf(stderr, "Error: min-size cannot be greater than max-size\n");
        return 1;
    }
    
    if (config.num_threads < 1) {
        fprintf(stderr, "Error: threads must be at least 1\n");
        return 1;
    }
    
    run_bubble(&config);
    
    return 0;
}
