#define _DEFAULT_SOURCE  

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <omp.h>
#include <unistd.h>
#include <time.h>
#include <signal.h>
#include <sys/time.h>

// LFSR random number generator thread local storage for OpenMP
_Thread_local unsigned lfsr = 1;
#define MASK 0xd0000001u
#define rand() (lfsr = (lfsr >> 1) ^ (unsigned int)((0 - (lfsr & 1u)) & MASK))

// Configuration structure
typedef struct {
    int min_working_set_mb;
    int max_working_set_mb;
    int num_threads;
    int runtime_seconds;
    double pressure_level; // 0.0 to 1.0
} bubble_config_t;

// Global variables
volatile int stop_flag = 0;
bubble_config_t config;

void signal_handler(int sig) {
    (void)sig; 
    stop_flag = 1;
    if (sig == SIGALRM) {
        printf("Runtime limit reached, stopping threads...\n");
    }
}

// Initialize LFSR with thread-specific seed
void init_thread_lfsr(int thread_id) {
    lfsr = 1 + thread_id; // Each thread gets a different seed
}

// Manual SSA memory operations for maximum parallelism
void random_memory_ops(double *data_chunk, int footprint_size, int thread_id) {
    int dump[100] = {0}; // Initialize dump array
    
    printf("Thread %d started with random access pattern, working set size: %d elements\n", 
           thread_id, footprint_size);
    
    while (!stop_flag) {
        unsigned r = rand() % footprint_size;
        
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
    }
    
    printf("Thread %d finished random access\n", thread_id);
}

// Streaming memory access pattern (based on STREAM benchmark)
void streaming_memory_ops(double *bw_data, int bw_stream_size, int thread_id) {
    double scalar = 3.0;
    
    printf("Thread %d started with streaming access pattern, working set size: %d elements\n", 
           thread_id, bw_stream_size);
    
    while (!stop_flag) {
        double *mid = bw_data + (bw_stream_size / 2);
        
        for (int i = 0; i < bw_stream_size / 2; i++) {
            bw_data[i] = scalar * mid[i];
        }

        for (int i = 0; i < bw_stream_size / 2; i++) {
            mid[i] = scalar * bw_data[i];
        }
    }
    
    printf("Thread %d finished streaming access\n", thread_id);
}

// Calculate working set size based on pressure level
int calculate_working_set_size(double pressure_level, int min_size, int max_size) {
    return min_size + (int)(pressure_level * (max_size - min_size));
}

// Load configuration from environment variables or use defaults
void load_config() {
    char *env_val;
    
    // Default configuration
    config.min_working_set_mb = 1;
    config.max_working_set_mb = 1024;
    config.num_threads = 4;
    config.runtime_seconds = 60;
    config.pressure_level = 0.5;
    

    if ((env_val = getenv("BUBBLE_MIN_WORKING_SET_MB")) != NULL) {
        config.min_working_set_mb = atoi(env_val);
    }
    if ((env_val = getenv("BUBBLE_MAX_WORKING_SET_MB")) != NULL) {
        config.max_working_set_mb = atoi(env_val);
    }
    if ((env_val = getenv("BUBBLE_NUM_THREADS")) != NULL) {
        config.num_threads = atoi(env_val);
    }
    if ((env_val = getenv("BUBBLE_RUNTIME_SECONDS")) != NULL) {
        config.runtime_seconds = atoi(env_val);
    }
    if ((env_val = getenv("BUBBLE_PRESSURE_LEVEL")) != NULL) {
        config.pressure_level = atof(env_val);
        if (config.pressure_level < 0.0) config.pressure_level = 0.0;
        if (config.pressure_level > 1.0) config.pressure_level = 1.0;
    }
    
    printf("Configuration:\n");
    printf("  Min working set: %d MB\n", config.min_working_set_mb);
    printf("  Max working set: %d MB\n", config.max_working_set_mb);
    printf("  Number of threads: %d\n", config.num_threads);
    printf("  Runtime: %d seconds\n", config.runtime_seconds);
    printf("  Pressure level: %.2f\n", config.pressure_level);
}

int main(int argc, char *argv[]) {

    load_config();
    
    if (argc > 1) {
        config.pressure_level = atof(argv[1]);
        if (config.pressure_level < 0.0) config.pressure_level = 0.0;
        if (config.pressure_level > 1.0) config.pressure_level = 1.0;
        printf("Pressure level overridden by command line: %.2f\n", config.pressure_level);
    }
    
    // Setup signal handling
    signal(SIGINT, signal_handler);
    signal(SIGTERM, signal_handler);
    signal(SIGALRM, signal_handler);
    
    omp_set_num_threads(config.num_threads);
    
    // Calculate actual working set size based on pressure level
    int working_set_size = calculate_working_set_size(
        config.pressure_level, 
        config.min_working_set_mb, 
        config.max_working_set_mb
    );
    
    printf("Starting bubble with working set size: %d MB (pressure: %.2f)\n", 
           working_set_size, config.pressure_level);
    printf("Using %d OpenMP threads\n", config.num_threads);
    
    // Calculate elements per working set
    int elements_per_working_set = working_set_size * 1024 * 1024 / sizeof(double);
    
    // Allocate shared memory for all threads
    size_t total_data_size = (size_t)config.num_threads * working_set_size * 1024 * 1024;
    double *shared_data = malloc(total_data_size);
    if (!shared_data) {
        fprintf(stderr, "Failed to allocate %zu bytes for shared data\n", total_data_size);
        return 1;
    }
    
    printf("Initializing %zu MB of memory...\n", total_data_size / (1024 * 1024));
    for (size_t i = 0; i < total_data_size / sizeof(double); i++) {
        shared_data[i] = (double)(i % 1000) + 1.0;
    }
    
    printf("Starting parallel execution...\n");
    
    if (config.runtime_seconds > 0) {
        alarm(config.runtime_seconds);
    }
    
    #pragma omp parallel
    {
        int thread_id = omp_get_thread_num();
        
        init_thread_lfsr(thread_id);
        
        double *thread_data = shared_data + (thread_id * elements_per_working_set);
        
        if (thread_id % 2 == 0) {
            random_memory_ops(thread_data, elements_per_working_set, thread_id);
        } else {
            streaming_memory_ops(thread_data, elements_per_working_set, thread_id);
        }
    }
    
    printf("All threads completed\n");
    
    free(shared_data);
    printf("Bubble execution completed\n");
    return 0;
}