#!/bin/bash

#Disable TLS certificate verification
kubectl config set-cluster minikube --insecure-skip-tls-verify=true

#Update minikube ip in config file
cat <<EOF > minikube/config
apiVersion: v1
clusters:
- cluster:
    extensions:
    - extension:
        provider: minikube.sigs.k8s.io
        version: v1.31.2
      name: cluster_info
    server: https://$(minikube ip):8443
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

#RBAC for anonymous user
kubectl apply -f minikube/rbac.yaml