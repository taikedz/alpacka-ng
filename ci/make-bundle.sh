#!/usr/bin/env bash

set -euo pipefail

HERE="$(dirname "$0")"
cd "$HERE/.."

require() {
    pkg="$1"
    if ! which $pkg; then
        echo "$pkg needed to produce build"
    fi
}

require podman
require tar
require go

current_ver="$(go run ci/version-is-bumped.go show)"

rm -rf ./bin/

bash build.sh
bash ci/build-alpine.sh

mkdir -p paf-bundle
mkdir -p bundles
cp -r bin/ paf-bundle/bin
cp ci/install-script.sh paf-bundle/install.sh
chmod 755 paf-bundle/install.sh

tar czf bundles/paf-"$current_ver"-bundle.tar.gz ./paf-bundle

rm -rf ./paf-bundle/