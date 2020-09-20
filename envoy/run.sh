#!/bin/bash

docker run --rm -p 10000:10000 -d --name envoy \
  -v $(pwd)/envoy.yaml:/etc/envoy/envoy.yaml \
  envoyproxy/envoy-dev:5d95032baa803f853e9120048b56c8be3dab4b0d
