#!/bin/sh
echo "generating swagger API docs for ${1}"

swag fmt
swag init -g services/raven/cmd/main.go -o ./services/raven/docs