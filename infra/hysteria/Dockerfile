FROM alpine:latest

RUN apk add --no-cache curl iproute2 iptables bash \
    && mkdir -p /etc/hysteria

# Download Hysteria2 binary
RUN curl -L -o /usr/local/bin/hysteria https://github.com/apernet/hysteria/releases/latest/download/hysteria-linux-amd64 \
    && chmod +x /usr/local/bin/hysteria

COPY config.yaml /etc/hysteria/config.yaml

CMD ["hysteria", "client", "-c", "/etc/hysteria/config.yaml"]
