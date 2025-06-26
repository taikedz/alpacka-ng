#!/usr/bin/env bash

set -euo pipefail

HERE="$(dirname "$0")"
cd "$HERE/.."


if [[ ! -f .venv/bin/activate ]]; then
    python3 -m venv .venv
    . .venv/bin/activate
    pip install -r ci/requirements.txt
else
    . .venv/bin/activate
fi

pytest test-behaviour/ "$@"
