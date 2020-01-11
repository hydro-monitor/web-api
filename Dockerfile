FROM alpine

ENV PORT 8080
ENV DB_HOSTS localhost
ENV DB_KEYSPACE hydromon
COPY _output/web-api /
ENTRYPOINT ./web-api
