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

rm -rf ./bin/

bash build.sh
bash ci/build-alpine.sh

mkdir -p paf-bundle
cp -r bin/ paf-bundle/bin
cp ci/install-script.sh paf-bundle/install.sh
chmod 755 paf-bundle/install.sh

tar czf paf-bundle.tar.gz ./paf-bundle

rm -rf ./paf-bundle/