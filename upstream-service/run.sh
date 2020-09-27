#!/bin/bash

[[ $# -ne 3 ]] && echo "please enter arguments" && exit 1
app=$1
num=$2
port=$3
image_name=upstream-service
container_name=${app}-${num}
network="envoy"

[[ `docker network ls | grep ${network} | wc -l` -ne 1 ]] && docker network create ${network}
[[ `docker ps -a -f name=${container_name} -q | wc -l` -eq 1 ]] && docker stop ${container_name} && docker rm -v ${container_name}
docker run -d \
  -p ${port}:8080 \
  --network ${network} \
  --name ${container_name} ${app}:latest