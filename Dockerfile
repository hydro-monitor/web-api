FROM ubuntu

WORKDIR /hydromon-server
COPY _output/web-api /hydromon-server
COPY configs/ /hydromon-server/scripts
COPY scripts/ /hydromon-server/scripts
COPY server.key /hydromon-server
COPY server.csr /hydromon-server

ENTRYPOINT ./web-api
