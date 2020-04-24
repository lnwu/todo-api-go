#!/bin/sh

set -euo pipefail

docker-compose up --force-recreate -d
docker image prune -f