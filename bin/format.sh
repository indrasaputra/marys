#!/usr/bin/env bash

set -euo pipefail

for file in `find . -name '*.go'`; do
  # Defensive, just in case.
  if [[ -f ${file} ]]; then
    awk '/^import \($/,/^\)$/{if($0=="")next}{print}' ${file} > /tmp/file
    mv /tmp/file ${file}
  fi
done

goimports -w -local github.com/indrasaputra/marys $(go list -f {{.Dir}} ./...)
gofmt -s -w .
