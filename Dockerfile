FROM alpine

ENV PORT 8080
ENV DB_HOSTS hydromon-cassandra-1,hydromon-cassandra-2,hydromon-cassandra-3
ENV DB_KEYSPACE hydromon
ENV DB_REPLICATION_FACTOR 3

WORKDIR /hydromon-server
COPY _output/web-api /hydromon-server
COPY scripts/ /hydromon-server/scripts

ENTRYPOINT ./web-api
