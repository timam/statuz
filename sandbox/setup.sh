#!/bin/bash

# Get the IP address of en0
ip_address=$(ifconfig en0 | awk '/inet / {print $2}')

# Check if IP address is not empty
if [ -n "$ip_address" ]; then
  echo "IP address of en0: $ip_address"

  # Disable TLS certificate verification
  kubectl config set-cluster minikube --insecure-skip-tls-verify=true

  # Update minikube IP in config file
  cat <<EOF > minikube/config
  apiVersion: v1
  clusters:
  - cluster:
      extensions:
      - extension:
          provider: minikube.sigs.k8s.io
          version: v1.31.2
        name: cluster_info
      server: http://$ip_address:18001  # Use en0 IP here
    name: minikube
  contexts:
  - context:
      cluster: minikube
      extensions:
      - extension:
          provider: minikube.sigs.k8s.io
          version: v1.31.2
        name: context_info
      namespace: default
      user: minikube
    name: minikube
  current-context: minikube
  kind: Config
  preferences: {}
  users:
  - name: minikube
EOF

  # RBAC for anonymous user
  kubectl apply -f minikube/rbac.yaml

else
  echo "IP address of en0 not found"
fi
