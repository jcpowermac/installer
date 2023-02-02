#!/bin/sh
# Example:  ./hack/go-test.sh

if [ "$IS_CONTAINER" != "" ]; then
  # there is no t.Short() in the repo
  go test -coverpkg=./cmd/...,./pkg/... -coverprofile=/tmp/coverprofile.out ./cmd/... ./data/... ./pkg/... "${@}"
  go tool cover -o cover.txt -func /tmp/coverprofile.out
  go tool cover -o cover.html -html /tmp/coverprofile.out
else
  podman run --rm \
    --env IS_CONTAINER=TRUE \
    --security-opt label=disable \
    --volume "${PWD}:/go/src/github.com/openshift/installer:rw" \
    --workdir /go/src/github.com/openshift/installer \
    docker.io/golang:1.18 \
    ./hack/go-test.sh "${@}"
fi
