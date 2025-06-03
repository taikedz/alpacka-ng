#!/usr/bin/env bash

HERE="$(dirname "$0")"
cd "$HERE"

go build -o bin/paf paf.go
