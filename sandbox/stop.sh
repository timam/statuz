#!/bin/bash

# Stop Docker Compose services
docker-compose down

# Find and stop the kubectl proxy process
kubectl_proxy_pid=$(ps aux | grep "kubectl proxy" | grep -v grep | awk '{print $2}')
if [ -n "$kubectl_proxy_pid" ]; then
  echo "Stopping kubectl proxy (PID: $kubectl_proxy_pid)..."
  kill "$kubectl_proxy_pid"
else
  echo "kubectl proxy process not found."
fi
