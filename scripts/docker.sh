#!/bin/bash

dir="$(dirname $(realpath ${BASH_SOURCE[0]}))" && cd $dir
dir=$dir/.. && cd $dir

docker buildx build \
--push \
--platform linux/arm64/v8,linux/amd64 \
--tag post:latest --force-rm \
-f Dockerfile .
if ! [ -z "$(docker images -f "dangling=true" -q 2>/dev/null)" ]
then
    docker rmi -f $(docker images -f "dangling=true" -q)
fi
docker images