.PHONY: all web-api

IMAGE_NAME=$(if $(ENV_IMAGE_NAME),$(ENV_IMAGE_NAME),hydro-monitor/web-api)
IMAGE_VERSION=$(if $(ENV_IMAGE_VERSION),$(ENV_IMAGE_VERSION),v0.0.0)

$(info web-api image settings: $(IMAGE_NAME) version $(IMAGE_VERSION))

all: web-api test

test:
	go test ./pkg/... -cover
	go vet ./pkg/...

web-api:
	go build -o _output/web-api ./cmd

web-api-linux:
	env GOOS=linux GOARCH=amd64 go build -o _output/web-api ./cmd
	docker build -t hydromon-server .

image-web-api:
	go mod vendor
	docker build -t $(IMAGE_NAME):$(IMAGE_VERSION) -f deploy/docker/Dockerfile .

push-image-web-api: image-web-api
	docker push $(IMAGE_NAME):$(IMAGE_VERSION)

clean:
	rm -f _output/web-api
	rm -f deploy/docker/web-api
