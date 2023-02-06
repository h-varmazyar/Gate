#!/bin/sh
echo "building ${1} with version ${2}"
docker build -f ./services/"${1}"/deploy/Dockerfile \
  --build-arg VERSION="${2}" \
  --build-arg GOPROXY="${GOPROXY}" \
  -t "${1}":"${2}" \
  .
