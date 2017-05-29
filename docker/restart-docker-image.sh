#!/usr/bin/env bash

#set -e

CONTAINER_NAME=$1
CONTAINER_ID=$(docker ps -a | awk '{print$1}' | xargs | awk '{print$2}')

if [ ${CONTAINER_ID} ]; then
  docker kill $CONTAINER_ID
  echo "Docker container $CONTAINER_ID has been killed"

  docker rm $CONTAINER_ID
  echo "Docker container $CONTAINER_ID has been removed"
else
  echo "There is no container running, ready to go!"
fi

echo "Preparing to run a new container..."
docker run -it --rm --name $CONTAINER_NAME $CONTAINER_NAME
