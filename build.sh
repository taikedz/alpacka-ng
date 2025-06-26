#!/usr/bin/env bash

BUILD=(go build -o bin/paf paf.go)
OCI=(run -i --rm -v "$PWD:/paf-build" -v "$PWD/gocache:/go" golang:1.24 sh -c "cd /paf-build; ${BUILD[*]}")


HERE="$(dirname "$0")"
cd "$HERE"

mkdir -p gocache

has() { which "$1" 2>&1 >/dev/null ; }
doit() { (set -x ; "$@") ; }

if has docker; then
    doit docker "${OCI[@]}"
elif has podman; then
    doit podman "${OCI[@]}"
fi
