#!/usr/bin/env bash
set -euxo pipefail

docker kill hydromon-cassandra-1
docker rm hydromon-cassandra-1
docker network rm hydromon-net
