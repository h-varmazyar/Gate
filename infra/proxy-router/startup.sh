#!/bin/bash

set -e

echo "[proxy-router] Starting proxy setup..."

# Create ipset set for routing via VPN
ipset create proxylist hash:ip || echo "ipset already exists"

# Configure DNS to use dnsmasq
echo "nameserver 127.0.0.1" > /etc/resolv.conf

# Start dnsmasq in background
dnsmasq --conf-file=/etc/dnsmasq.conf &
echo "[proxy-router] dnsmasq started..."

# Wait for dnsmasq to populate ipset
sleep 3

# Setup iptables marking for matched IPs
iptables -t mangle -A OUTPUT -m set --match-set proxylist dst -j MARK --set-mark 1

# Setup policy routing to use hysteria tunnel for marked packets
ip rule add fwmark 1 table 100
ip route add default dev utun table 100

# Optional NAT (if needed)
iptables -t nat -A POSTROUTING -o utun -j MASQUERADE

echo "[proxy-router] Routing rules applied. Container is ready."

# Keep container running
tail -f /dev/null
