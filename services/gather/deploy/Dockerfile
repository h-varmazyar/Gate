FROM golang:alpine as builder

ARG VERSION
ARG BUILD_ID

LABEL stage=builder
LABEL build=$BUILD_ID

WORKDIR /src
COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 go build -buildvcs=false -a -installsuffix cgo \
    -ldflags "-w -X main.Version=${VERSION}" \
    -o ./build/gather ./services/gather/cmd

# ─────────────────────────────────────────────

FROM alpine

WORKDIR /app
COPY --from=builder /src/build/ ./
#COPY ./services/gather/ips ./ips
COPY ./services/gather/migrations ../migrations
ENTRYPOINT [ "./gather" ]