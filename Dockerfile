# Build status-go in a Go builder container
FROM golang:1.10-alpine as builder

ARG build_tags
ARG build_flags

RUN apk add --no-cache git make gcc musl-dev linux-headers

ENV REPO=/go/src/github.com/status-im/p2p-health-bot

RUN mkdir -p ${REPO}
ADD . ${REPO}
RUN cd ${REPO} && go get && go build

# Copy the binary to the second image
FROM alpine:latest

RUN apk add --no-cache ca-certificates bash

ENV REPO=/go/src/github.com/status-im/p2p-health-bot

COPY --from=builder ${REPO}/p2p-health-bot /usr/local/bin/

EXPOSE 8080
