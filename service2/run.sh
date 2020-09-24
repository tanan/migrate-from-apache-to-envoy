#!/bin/bash

[[ $# -ne 1 ]] && echo "please enter an argument" && exit 1
num=$1
app=service2
container_name=${app}-${num}
network="envoy"

[[ `docker network ls | grep ${network} | wc -l` -ne 1 ]] && docker network create ${network}
[[ `docker ps -a -f name=${container_name} -q | wc -l` -eq 1 ]] && docker stop ${container_name} && docker rm -v ${container_name}
docker run -d \
  --network ${network} \
  --name ${container_name} ${app}:latest