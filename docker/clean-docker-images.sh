!#/usr/bin/env bash

readonly HAS_GARBAGE=$(docker images -f "dangling=true" -q)

if [[ ! -z $HAS_GARBAGE ]]; then
  docker rmi -f $HAS_GARBAGE
fi
