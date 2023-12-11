cleanup:
  rm -rf ./ui/dist
  rm -rf ./ui/.next
  rm -rf ./bin

build-binary:
  #!/usr/bin/env bash
  set -euxo pipefail

  just cleanup
  echo "Building UI..."
  pushd ui
  yarn install
  yarn export
  popd
  echo "Building server..."
  go build -o bin/golang-aio main.go
  echo "Built artifact at ./bin/golang-aio"

build-docker:
  @echo "Building docker image..."
  docker build -t golang-aio .
