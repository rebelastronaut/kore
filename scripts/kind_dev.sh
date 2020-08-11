#!/bin/sh

set -eo pipefail

usage() {
cat <<EOF
Usage: $(basename $0)
  --recreate    if set the kind cluster will be recreated
  -h|--help     display this usage menu
EOF
  if [[ -n $@ ]]; then
    echo "[error] $@"
    exit 1
  fi
  exit 0
}

recreate=false

while [[ $# -gt 0 ]]; do
  case "$1" in
  --recreate)    recreate=true;    shift 1; ;;
  -h|--help)     usage;            ;;
  *)                               shift 1; ;;
  esac
done

PROJECT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )"/.. >/dev/null 2>&1 && pwd )"

if [ "$recreate" = "true" ]; then
  kind delete cluster --name kore
fi

if ! kind get clusters | grep "kore" ; then
cat <<EOF | kind create cluster --name kore --config=-
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
  - role: control-plane
    image: kindest/node:v1.15.11@sha256:6cc31f3533deb138792db2c7d1ffc36f7456a06f1db5556ad3b6927641016f50
    extraPortMappings:
      - containerPort: 10080
        hostPort: 10080
        protocol: TCP
      - containerPort: 10443
        hostPort: 10443
        protocol: TCP
    extraMounts:
      - hostPath: $PROJECT_DIR
        containerPath: /go/src/github.com/appvia/kore
EOF
fi

kubectl config use-context kind-kore

if ! kubectl get ns kore; then
  kubectl create ns kore
fi

make -C "$PROJECT_DIR" kind-image-dev

helm upgrade -i --namespace kore kore "$PROJECT_DIR/charts/kore" --wait -f "$PROJECT_DIR/charts/my_values.yaml"
