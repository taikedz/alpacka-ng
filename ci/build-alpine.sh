#!/usr/bin/env bash

set -euo pipefail

HERE="$(dirname "$0")"
cd "$HERE"

podman build -t paf-build:alpine ./alpine

podman run -v "$PWD/..:/hostdata" -it --rm paf-build:alpine sh -c "cd /hostdata && go build -o bin/paf-alpine paf.go"
