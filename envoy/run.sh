#!/bin/bash

name="envoy"
network="envoy"

[[ `docker network ls | grep ${network} | wc -l` -ne 1 ]] && docker network create ${network}
[[ `docker ps -a -f name=${name} -q | wc -l` -eq 1 ]] && docker stop ${name} && docker rm -v ${name}
docker run -p 80:80 -d --name envoy \
  -v $(pwd)/envoy.yaml:/etc/envoy/envoy.yaml \
  -e "ENVOY_UID=0" \
  --network ${network} \
  envoyproxy/envoy-dev:5d95032baa803f853e9120048b56c8be3dab4b0d
