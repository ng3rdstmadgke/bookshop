#!/bin/bash


function usage {
cat >&2 << EOF
Usage: $0 [OPTIONS]

Build Docker images for the project.

Options:
  -h, --help        Show this help message and exit
EOF
exit 1
}

source $PROJECT_DIR/bin/lib/utils.sh


PUSH=false
DOCKER_USER=keitamido
VERSION=latest
args=()
while [[ "$#" -gt 0 ]]; do
  case $1 in
    -h | --help       ) usage; exit 0 ;;
    -v | --version    ) shift; VERSION="$1" ;;
    -p | --push       ) PUSH=true ;;
    --docker-user      ) shift; DOCKER_USER="$1" ;;
    *                 ) args+=("$1") ;;
  esac
  shift
done


if [ ${#args[@]} -ne 0 ]; then
  usage
fi

set -eo pipefail

for component in catalogue bff frontend; do
  IMAGE_NAME="bookshop-$component"
  info docker build -t $DOCKER_USER/$IMAGE_NAME:$VERSION $PROJECT_DIR/application/$component
  docker build -t $DOCKER_USER/$IMAGE_NAME:$VERSION $PROJECT_DIR/application/$component
done


if [ "$PUSH" = "true" ]; then
  docker login
  for component in catalogue bff frontend; do
    IMAGE_NAME="bookshop-$component"
    info docker push $DOCKER_USER/$IMAGE_NAME:$VERSION
    docker push $DOCKER_USER/$IMAGE_NAME:$VERSION
  done
fi