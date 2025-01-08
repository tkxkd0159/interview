#!/usr/bin/env bash

# apiserver proxy URL format:
# https://:kubernetes_master_address/api/v1/namespaces/:namespace_name/services/[https:]service_name[:port_name]/proxy

# Get the API server URL from kubeconfig
API_SERVER=$(kubectl config view --minify -o jsonpath='{.clusters[0].cluster.server}')

# Get the token of the service account
TOKEN=$(kubectl get secret/test-client-ac -o jsonpath='{.data.token}' | base64 -d)

# Set k8s cluster X.509 certificate (cluster pubkey)
CA_CERT_PATH=/home/user/.minikube/ca.crt

NAMESPACE=default
RESOURCE_PATH=api/v1/namespaces/$NAMESPACE/services
DEST=$API_SERVER/$RESOURCE_PATH/nginx-service/proxy/

# Access the custom resource via kube-apiserver
curl -L "$DEST" -H "Authorization: Bearer $TOKEN" \
     --cacert $CA_CERT_PATH
