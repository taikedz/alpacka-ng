#!/usr/bin/env bash

HERE="$(dirname "$0")"
cd "$HERE"

go build -trimpath -o bin/paf paf.go

