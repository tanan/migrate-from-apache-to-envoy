#!/bin/bash

name=upstream-app
network="envoy"

[[ `docker network ls | grep ${network} | wc -l` -ne 1 ]] && docker network create ${network}
[[ `docker ps -a -f name=${name} -q | wc -l` -eq 1 ]] && docker stop ${name} && docker rm -v ${name}
docker run --rm -d \
  -p 8080:8080 \
  --network ${network} \
  --name ${name} ${name}:latest