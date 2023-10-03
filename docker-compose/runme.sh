#!/bin/bash
kubectl config set-cluster minikube --insecure-skip-tls-verify=true

 kubectl config view --minify --raw --flatten -o jsonpath='{.clusters[].cluster.certificate-authority-data}' | base64 --decode
 kubectl config view --minify --raw --flatten -o jsonpath='{.users[].user.client-key-data}' | base64 --decode

# Copy client.crt and client.key to the minikube directory
cp ~/.minikube/profiles/minikube/client.key minikube/
cp ~/.minikube/profiles/minikube/client.crt minikube/