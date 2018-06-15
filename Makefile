COMMIT = $(shell git rev-parse --short HEAD)
IMAGE_NAME = 'statusteam/p2p-health-bot'

all: docker-image

docker-image:
	docker build \
		-t $(IMAGE_NAME):$(COMMIT) \
		-t $(IMAGE_NAME):latest \
		.
