```sh
export SERVER_IP=10.10.1.2
sudo apt install apache2-utils
```

# Read
```sh
ab -n 99999999999 -c 100 -k http://$SERVER_IP:80/
```

# Write
```sh
ab -n 99999999999 -c 100 -p /dev/urandom -T application/octet-stream http://$SERVER_IP:80/
```