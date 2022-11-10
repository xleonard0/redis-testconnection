# Redis Test Connection Tool 🚀

Simple tool written in Go to test connection to Redis server with authentication enabled.

## Environment variables


> **REDIS_SERVER :** Redis IP + Port (Example: localhost:6379)

> **REDIS_PASSWORD** 

## How to use?

[Download the tool](https://github.com/xleonard0/redis-testconnection/releases)

### Linux 

``` bash
REDIS_SERVER="REDIS_IP:REDIS_PORT" REDIS_PASSWORD="PASSWORD!" ./redis-testconnection-linuxamd64 
```

### Windows 

``` cmd
set "REDIS_SERVER=REDIS_IP:REDIS_PORT" && set "REDIS_PASSWORD=PASSWORD" && redis-testconnection-windowsamd64.exe
```
