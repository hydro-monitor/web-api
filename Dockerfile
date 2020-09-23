FROM ubuntu

WORKDIR /hydromon-server
COPY _output/web-api /hydromon-server
COPY configs/ /hydromon-server/configs
COPY scripts/ /hydromon-server/scripts
COPY server.key /hydromon-server
COPY server.csr /hydromon-server

ENTRYPOINT ./web-api
