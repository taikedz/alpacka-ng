#!/usr/bin/env bash

set -euo pipefail

HERE="$(dirname "$0")"

if ! go run "$HERE/version-is-bumped.go"; then
    exit 1
fi

bash "$HERE/../build.sh"

if ! git status | grep "working tree clean" -q; then
    echo "Please ensure clean working tree before proceeding."
    exit 1
fi

newtag="$(go run "$HERE/version-is-bumped.go" show)"
git tag "$newtag"
git push
git push --tags

echo "Done"