FROM alpine

ENV PORT 8080
ENV DB_HOSTS hydromon-cassandra-1
ENV DB_KEYSPACE hydromon

WORKDIR /hydromon-server
COPY _output/web-api /hydromon-server
COPY scripts/ /hydromon-server/scripts

ENTRYPOINT ./web-api
