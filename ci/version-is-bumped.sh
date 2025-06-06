#!/usr/bin/env bash

latest_tag="$(git log --oneline --decorate=short --color=never | grep -Po 'tag: v\d+\.\d+\.\d+'|sed 's/tag: v//'|head -n 1)"

if [[ -z "$latest_tag" ]]; then
    exit 0
fi

current_ver="$(grep -Po 'v\d+\.\d+\.\d+' pakang/version.go|sed 's/v//')"

python3 -c "p = lambda s: [int(v) for v in s.split('.')] ; exit(0) if p('$latest_tag') < p('$current_ver') else exit(1)"
exit $res 