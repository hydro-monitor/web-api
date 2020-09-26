FROM ubuntu

WORKDIR /hydromon-server
COPY _output/web-api /hydromon-server
COPY configs/ /hydromon-server/configs
COPY scripts/ /hydromon-server/scripts
COPY fullchain.pem /hydromon-server
COPY privkey.pem /hydromon-server

ENTRYPOINT ./web-api
