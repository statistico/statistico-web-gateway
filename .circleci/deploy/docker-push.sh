#!/bin/bash

set -e

aws ecr get-login --no-include-email --region $AWS_DEFAULT_REGION | bash

docker tag "statisticowebgateway_rest" "$AWS_ECR_ACCOUNT_URL/statistico-web-gateway:$CIRCLE_SHA1"
docker push "$AWS_ECR_ACCOUNT_URL/statistico-web-gateway:$CIRCLE_SHA1"
