#!/usr/bin/env bash

set -euo pipefail

PROJECT_MODULE="github.com/traefik/hub-crds"
IMAGE_NAME="kubernetes-codegen:latest"
CURRENT_DIR="$(pwd)"

echo "Building codegen Docker image..."
docker build --build-arg KUBE_VERSION=v0.20.15 -f "./script/codegen.Dockerfile" \
  --build-arg USER="${USER}" \
  --build-arg UID="$(id -u)" \
  --build-arg GID="$(id -g)" \
  -t "${IMAGE_NAME}" \
  "."

cmd="/go/src/k8s.io/code-generator/generate-groups.sh all $PROJECT_MODULE/hub/v1alpha1/generated $PROJECT_MODULE hub:v1alpha1"

echo "Generating Hub clientset, listers and informers code ..."
docker run --rm \
  -v "${CURRENT_DIR}:/go/src/${PROJECT_MODULE}" \
  -w "/go/src/${PROJECT_MODULE}" \
  "${IMAGE_NAME}" $cmd

cmd="controller-gen crd:crdVersions=v1 paths=./hub/v1alpha1/... output:crd:dir=./hub/v1alpha1/crd"

echo "Generating the CRD definitions ..."
docker run --rm \
  -v "${CURRENT_DIR}:/go/src/${PROJECT_MODULE}" \
  -w "/go/src/${PROJECT_MODULE}" \
  "${IMAGE_NAME}" $cmd
