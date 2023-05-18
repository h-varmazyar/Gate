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
RUN swag init -g ./services/raven/cmd/main.go -o ./services/raven/docs

RUN GO111MODULE=on CGO_ENABLED=0 go build -buildvcs=false -a -installsuffix cgo \
    -ldflags "-w -X main.Version=${VERSION}" \
    -o ./build/raven ./services/raven/cmd

FROM alpine
WORKDIR /app
COPY --from=builder /src/build/ ./
COPY ./services/raven/docs ./docs
ENTRYPOINT [ "./raven" ]

EXPOSE 8080