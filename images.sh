#!/bin/bash

TAG=$(date +%Y%m%d%H%M%S)
docker build -t pi-img:$TAG . --no-cache

if [ "$(docker ps -a --format "{{.Names}}" | grep pi-cont)" = "pi-cont" ]; then
    docker stop pi-cont
    docker rm pi-cont
fi

docker create --name pi-cont pi-img:$TAG
docker cp pi-cont:/go/bin/pi ./pi
docker rm pi-cont

docker tag pi-img:$TAG pi-img:latest
docker run --name pi-cont -d pi-img:latest pi -h 8080


IMAGES=$(docker images pi-img --format "{{.Tag}}" | wc -l)

if (( $IMAGES > 0 )); then
    docker rmi -f $(docker images pi-img --format "{{.Repository}}:{{.Tag}}" | tail -1)
fi