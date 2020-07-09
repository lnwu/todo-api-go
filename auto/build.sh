#!/bin/sh

set -eu

TAG=ccr.ccs.tencentyun.com/k8-test/todo-api:${CI_BUILD_NUMBER}

docker build . -t ${TAG}

docker login -u ${DOCKER_USERNAME} -p ${DOCKER_PASSWORD} ccr.ccs.tencentyun.com
docker push ${TAG}
docker logout ccr.ccs.tencentyun.com