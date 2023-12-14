cleanup:
  rm -rf ./web/dist
  rm -rf ./web/.next
  rm -rf ./bin

build-binary:
  #!/usr/bin/env bash
  set -euxo pipefail

  just cleanup
  echo "Building Web..."
  pushd web
  yarn install
  yarn export
  popd
  echo "Building server..."
  go build -o bin/golang-aio main.go
  echo "Built artifact at ./bin/golang-aio"

build-docker:
  @echo "Building docker image..."
  docker build -t golang-aio .
