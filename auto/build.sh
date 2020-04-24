#!/bin/sh

set -euo pipefail

docker build . -t lnwu/todo-api:${BUILDKITE_BUILD_NUMBER}