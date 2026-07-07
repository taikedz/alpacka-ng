#!/usr/bin/env bash

THIS="$(readlink -f "$0")"
HEREDIR="$(dirname "$THIS")"

set -euo pipefail

cd "$HEREDIR"

if [[ -z "${1:-}" ]]; then
    version="$(git tag|sort -V|tail -n1)"
else
    version="$1"
fi

echo "--- Downloading version $version"

wget -q https://github.com/taikedz/alpacka-ng/releases/download/${version}/paf || {
    res=$?
    echo "FATAL : Could not download version '$version'"
    exit $res
}
chmod 755 ./paf

sudo mv ./paf /usr/local/bin/paf

echo "Installed."
