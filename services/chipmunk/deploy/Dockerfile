FROM golang:alpine as builder
ARG VERSION
WORKDIR /src
COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 go build -buildvcs=false -a -installsuffix cgo \
    -ldflags "-w -X main.Version=${VERSION}" \
    -o ./build/chipmunk ./services/chipmunk/cmd

FROM alpine
WORKDIR /app
COPY --from=builder /src/build/ ./
COPY ./services/chipmunk/assets ./assets
ENTRYPOINT [ "./chipmunk" ]