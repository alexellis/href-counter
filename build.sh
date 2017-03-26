#!/bin/sh
echo Building alexellis2/href-counter:build

docker build -t alexellis2/href-counter:build . -f Dockerfile.build

docker create --name extract alexellis2/href-counter:build 
docker cp extract:/go/src/github.com/alexellis/href-counter/app ./app
docker rm -f extract

echo Building alexellis2/href-counter:latest

docker build --no-cache -t alexellis2/href-counter:latest .
