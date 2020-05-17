#!/usr/bin/env bash
set -euxo pipefail

docker network create hydromon-net
docker run --name hydromon-cassandra-1 --net hydromon-net -d cassandra
sleep 15
docker run --network hydromon-net -v $(pwd):/tmp/mount --rm cassandra cqlsh hydromon-cassandra-1 -f "/tmp/mount/create-keyspace.cql" 
make web-api-linux
docker build -t hydromon-server .
docker run --net hydromon-net -p 8080:8080 -d hydromon-server

