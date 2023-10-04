#!/bin/bash

# Start kubectl proxy in the background
kubectl proxy --address='0.0.0.0' --port='18001' --disable-filter=true &

# Start Docker Compose services
docker-compose up -d
