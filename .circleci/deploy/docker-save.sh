#!/bin/bash

set -e

mkdir -p /tmp/workspace/docker-cache

docker save -o /tmp/workspace/docker-cache/statisticowebgateway_rest.tar statisticowebgateway_rest:latest
