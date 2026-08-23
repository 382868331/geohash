#!/usr/bin/env bash
set -euo pipefail
docker build -f goletalab.Dockerfile -t geohash-task .
docker run --rm geohash-task "$@"
