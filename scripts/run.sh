#!/bin/sh
echo "running ${1} ${2}"
cd ./services/"${1}"/cmd && go run . "${2}"