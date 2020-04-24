#!/bin/sh

set -euo pipefail

docker tag lnwu/todo-api:${BUILDKITE_BUILD_NUMBER} lnwu/todo-api-${STAGE}:${BUILDKITE_BUILD_NUMBER}
docker-compose -f ./auto/${STAGE}/docker-compose.yml up --force-recreate -d 
docker image prune -f -a