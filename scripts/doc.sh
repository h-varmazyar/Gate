#!/bin/sh
echo "generating swagger API docs for ${1}"

swag fmt
swag init -g services/gateway/cmd/main.go -o ./services/gateway/docs