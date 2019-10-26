#!/bin/sh
echo Building using multi-stage build

docker build --no-cache -t alexellis2/href-counter:0.1.0 . -f Dockerfile.multi
