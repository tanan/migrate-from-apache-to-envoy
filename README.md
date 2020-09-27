# migrate-from-apache-to-envoy

## try front envoy proxy example

- start upstream service containers

```:sh
$ cd upstreamp-service
$ docker build -t upstream-service:latest .
$ ./run.sh service1 1 8080
$ ./run.sh service1 2 8081
$ ./run.sh service2 1 8082
$ ./run.sh service2 2 8083
$ docker ps
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS                    NAMES
13db068a396c        service2:latest     "/upstream-app"     4 seconds ago       Up 3 seconds        0.0.0.0:8083->8080/tcp   service2-2
3b9224bcf50b        service2:latest     "/upstream-app"     7 seconds ago       Up 6 seconds        0.0.0.0:8082->8080/tcp   service2-1
79b51476b24f        service1:latest     "/upstream-app"     10 seconds ago      Up 10 seconds       0.0.0.0:8081->8080/tcp   service1-2
8fe6456e04c5        service1:latest     "/upstream-app"     15 seconds ago      Up 14 seconds       0.0.0.0:8080->8080/tcp   service1-1
```

- start an envoy container

```:sh
$ cd envoy/certs
$ openssl req -nodes -new -x509 \
  -keyout example-com.key -out example-com.crt \
  -days 365 \
  -subj '/CN=example.com/O=My Company Name LTD./C=US';
$ cd ..
$ ./run.sh
$ docker ps -f name=envoy
CONTAINER ID        IMAGE                                                           COMMAND                  CREATED             STATUS              PORTS                                                 NAMES
59accf835e00        envoyproxy/envoy-dev:5d95032baa803f853e9120048b56c8be3dab4b0d   "/docker-entrypoint.…"   14 seconds ago      Up 13 seconds       0.0.0.0:80->80/tcp, 0.0.0.0:443->443/tcp, 10000/tcp   envoy
```

- check request

```:sh
$ curl -k -L --resolve www.example.com:80:127.0.0.1 --resolve www.example.com:443:127.0.0.1 http://www.example.com/healthz
{"Host":"8fe6456e04c5","Service":"service1","Message":"Status is healthy"}
```

## e2e test

```:sh
$ cd e2e
$ go test ./specs
ok      e2e/specs       10.121s
```
