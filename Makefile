all: docker-image

docker-image:
	docker build -t statusteam/p2p-health-bot .
