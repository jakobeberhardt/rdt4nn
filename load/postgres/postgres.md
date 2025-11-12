```sh
export PGPASSWORD=benchmark
export SERVER_IP=10.10.1.2
```

## Large Dataset
mixed
```sh
pgbench -i -s 30 -h $SERVER_IP -U benchmark benchmark && pgbench -c 16 -j 16 -T 3600 -h $SERVER_IP -U benchmark benchmark
```

read
```sh
pgbench -i -s 30 -h $SERVER_IP -U benchmark benchmark && pgbench -c 16 -j 16 -T 3600 -S -h $SERVER_IP -U benchmark benchmark
```

write
```sh
pgbench -i -s 30 -h $SERVER_IP -U benchmark benchmark && pgbench -c 16 -j 16 -T 3600 -N -h $SERVER_IP -U benchmark benchmark
```

## Small Dataset
mixed
```sh
pgbench -i -s 1 -h $SERVER_IP -U benchmark benchmark && pgbench -c 16 -j 16 -T 3600 -h $SERVER_IP -U benchmark benchmark
```

read
```sh
pgbench -i -s 1 -h $SERVER_IP -U benchmark benchmark && pgbench -c 16 -j 16 -T 3600 -S -h $SERVER_IP -U benchmark benchmark
```

write
```sh
pgbench -i -s 1 -h $SERVER_IP -U benchmark benchmark && pgbench -c 16 -j 16 -T 3600 -N -h $SERVER_IP -U benchmark benchmark
```