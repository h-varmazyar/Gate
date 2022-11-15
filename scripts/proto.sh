#!/bin/sh
echo "generating ${1} protobuf codes"
if [ "${1}" = "" ]; then
  protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./api/proto/src/*.proto
  rm -rf ./api/proto/*.pb.go
  mv ./api/proto/src/*.pb.go ./api/proto/
  for f in ./api/proto/*.pb.go; do
    protoc-go-inject-tag -input="$f" -XXX_skip=json,xml,yaml
  done
else
  protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=require_unimplemented_servers=false:. --go-grpc_opt=paths=source_relative \
    ./services/"${1}"/api/proto/src/*.proto
  rm -rf ./services/"${1}"/api/proto/*.pb.go
  mv ./services/"${1}"/api/proto/src/*.pb.go ./services/"${1}"/api/proto/
  for f in ./services/"${1}"/api/proto/*.pb.go; do
    protoc-go-inject-tag -input="$f" -XXX_skip=json,xml,yaml
  done
fi
