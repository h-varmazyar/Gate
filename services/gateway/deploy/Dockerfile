FROM golang:alpine as builder

ARG VERSION
ARG BUILD_ID

LABEL stage=builder
LABEL build=$BUILD_ID

#todo: after build in CI/CD script run next command to remove intermediate image
# docker image prune --filter label=stage=builder --filter label=build=$BUILD_ID

WORKDIR /src
COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@v1.8.12
RUN swag init -g ./services/gateway/cmd/main.go -o ./services/gateway/docs

RUN GO111MODULE=on CGO_ENABLED=0 go build -buildvcs=false -a -installsuffix cgo \
    -ldflags "-w -X main.Version=${VERSION}" \
    -o ./build/gateway ./services/gateway/cmd

FROM alpine
WORKDIR /app
COPY --from=builder /src/build/ ./
COPY ./services/gateway/docs ./docs
ENTRYPOINT [ "./gateway" ]

EXPOSE 8080