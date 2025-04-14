#!/bin/bash

set -e

NETWORK_NAME="gate"
SUBNET="10.89.0.0/16"
GATEWAY="10.89.0.1"

echo "🔍 Checking for existing Docker network: $NETWORK_NAME"

if docker network inspect "$NETWORK_NAME" >/dev/null 2>&1; then
  echo "❌ Network '$NETWORK_NAME' already exists. Removing it..."
  docker network rm "$NETWORK_NAME"
else
  echo "✅ Network '$NETWORK_NAME' does not exist. Proceeding..."
fi

echo "🚀 Creating network '$NETWORK_NAME' with subnet $SUBNET"

if docker network create \
  --driver=bridge \
  --subnet="$SUBNET" \
  --gateway="$GATEWAY" \
  "$NETWORK_NAME"; then
  echo "✅ Network '$NETWORK_NAME' created successfully!"
else
  echo "❌ Failed to create network. The subnet may conflict with another existing network."
  exit 1
fi
