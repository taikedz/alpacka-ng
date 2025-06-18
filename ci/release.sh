#!/usr/bin/env bash

set -euo pipefail

HERE="$(readlink -f "$(dirname "$0")")"
cd "$HERE/.."

if ! git status | grep "working tree clean" -q; then
    echo "Please ensure clean working tree before proceeding."
    exit 1
fi

if ! go run "$HERE/version-is-bumped.go"; then
    exit 1
fi

bash build.sh

newtag="$(go run "$HERE/version-is-bumped.go" show)"
git tag "$newtag"
git push
git push --tags

echo "Done"