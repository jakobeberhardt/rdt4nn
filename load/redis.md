```sh
export SERVER_IP=10.10.1.2
export REDIS_SIZE=1024
```

# Read/Write
```sh
redis-benchmark -h $SERVER_IP  -t set,get -d $REDIS_SIZE -r 1000000 -n 99999999 -c 50 -P 16
```

# Read
```sh
redis-benchmark -h $SERVER_IP -t get -d $REDIS_SIZE -r 1000000 -n 99999999 -c 50 -P 16
```

# Write
```sh
redis-benchmark -h $SERVER_IP -t get -d $REDIS_SIZE -r 1000000 -n 99999999 -c 50 -P 16
```