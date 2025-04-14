FROM alpine:latest

RUN apk update && apk add --no-cache \
    iptables \
    iproute2 \
    ipset \
    dnsmasq \
    bash \
    curl \
    bind-tools

# Copy startup scripts
COPY ./infra/proxy-router/startup.sh /startup.sh
COPY ./infra/proxy-router/dnsmasq.conf /etc/dnsmasq.conf

# Make startup.sh executable
RUN chmod +x /startup.sh

# Run startup
CMD ["/startup.sh"]
