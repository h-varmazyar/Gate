FROM golang:alpine as builder
ARG VERSION
ARG GOPROXY
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 go build -a -installsuffix cgo \
    -ldflags "-w -X main.Version=${VERSION}" \
    -i -o ./build .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /src/build/ ./
COPY ./services/scms/assets ./assets
ENTRYPOINT [ "./gate" ]
