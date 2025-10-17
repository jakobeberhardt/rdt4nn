#!/usr/bin/env python3
import redis
import random
import string
import sys

# Connection parameters
REDIS_HOST = 'localhost'
REDIS_PORT = 6379
REDIS_DB = 0

KEY_SIZE = 64
TARGET_SIZE_GB = 1.0
TARGET_SIZE_BYTES = int(TARGET_SIZE_GB * 1024 * 1024 * 1024)

REDIS_OVERHEAD_PER_KV = 96

def generate_random_string(length):
    """Generate a random alphanumeric string of given length."""
    return ''.join(random.choices(string.ascii_letters + string.digits, k=length))

def format_bytes(bytes_val):
    """Format bytes into human-readable format."""
    for unit in ['B', 'KB', 'MB', 'GB']:
        if bytes_val < 1024.0:
            return f"{bytes_val:.2f} {unit}"
        bytes_val /= 1024.0
    return f"{bytes_val:.2f} TB"

def main():
    print("=" * 70)
    print("Redis Data Generator - 64-byte keys, ~1GB target")
    print("=" * 70)
    
    # Connect to Redis
    try:
        r = redis.Redis(host=REDIS_HOST, port=REDIS_PORT, db=REDIS_DB, 
                       socket_connect_timeout=5, decode_responses=True)
        r.ping()
        print(f"Connected to Redis at {REDIS_HOST}:{REDIS_PORT}")
    except Exception as e:
        print(f"Failed to connect to Redis: {e}")
        sys.exit(1)
    
    # Clear existing data
    r.flushdb()
    print("Cleared existing database")

    AVG_VALUE_SIZE = 1024
    estimated_keys = TARGET_SIZE_BYTES // (KEY_SIZE + AVG_VALUE_SIZE + REDIS_OVERHEAD_PER_KV)
    
    print(f"\nTarget size: {format_bytes(TARGET_SIZE_BYTES)}")
    print(f"Key size: {KEY_SIZE} bytes")
    print(f"Average value size: {AVG_VALUE_SIZE} bytes")
    print(f"Estimated keys needed: ~{estimated_keys:,}")
    print("\nGenerating data...")
    
    # Generate data in batches
    batch_size = 10000
    total_keys = 0
    current_size = 0
    
    # Use pipeline for better performance
    pipe = r.pipeline(transaction=False)
    
    while current_size < TARGET_SIZE_BYTES:
        for i in range(batch_size):
            # Generate key with exactly 64 bytes
            key_suffix_length = KEY_SIZE - 4  # 4 chars for "key_"
            key = f"key_{generate_random_string(key_suffix_length)}"
            
            # Generate value with variable size (512 to 2048 bytes)
            value_size = random.randint(512, 2048)
            value = generate_random_string(value_size)
            
            pipe.set(key, value)
            
            # Estimate size
            current_size += KEY_SIZE + value_size + REDIS_OVERHEAD_PER_KV
            total_keys += 1
            
            if current_size >= TARGET_SIZE_BYTES:
                break
        
        # Execute batch
        pipe.execute()
        
        # Progress update
        progress = (current_size / TARGET_SIZE_BYTES) * 100
        print(f"Progress: {progress:.1f}% | Keys: {total_keys:,} | "
              f"Est. size: {format_bytes(current_size)}", end='\r')
        
        if current_size >= TARGET_SIZE_BYTES:
            break
    
    print()
    
    # Get actual memory usage from Redis
    info = r.info('memory')
    actual_memory = info['used_memory']
    
    print("\n" + "=" * 70)
    print("Data Generation Complete!")
    print("=" * 70)
    print(f"Total keys inserted: {total_keys:,}")
    print(f"Estimated size: {format_bytes(current_size)}")
    print(f"Actual Redis memory: {format_bytes(actual_memory)}")
    print(f"Redis overhead: {format_bytes(actual_memory - current_size)}")
    print("=" * 70)
    
    print("\nSample keys (first 5):")
    keys = r.keys('key_*')[:5]
    for key in keys:
        value_len = len(r.get(key))
        print(f"  {key} -> value length: {value_len} bytes")

if __name__ == '__main__':
    main()
