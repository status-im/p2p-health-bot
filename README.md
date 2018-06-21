# P2P Health checking bot
---
This bot measures latency for Status messages (complete roundtrip).

It operates in two modes: sender and receiver. Sends sends health check messages periodically, and tracks responses from receiver bot.

# Installation

```
go get github.com/status-im/p2p-health-bot
```

You may also want to run `dep ensure` to use deps versions used at the moment of writing package.

# Usage

### First node
Start statusd:

```
./statusd -shh=true -standalone=false -http=true -status=http -networkid=1
```

Then, start bot in a sender mode:

```
p2p-health-bot
```


### Second node
Start statusd (make sure, it's the same network):

```
./statusd -shh=true -standalone=false -http=true -status=http -networkid=1
```

Then, start bot in a receiver mode:

```
p2p-health-bot -send=false
```

See `p2p-health-bot --help` for more options.

### Metrics

Metrics are exposed in Prometheus format on `/metrics` endpoint of **sender** node only. Default listen port is 8008. Use `-statsPort` flag to change it.

Currently exposed metrics:

 - `msgs_sent`
 - `msgs_received`
 - `msgs_responses_latency`

### Deployment

This software is deployed via Docker image built with `Dockerfile` ran from the `Makefile`. The image is then pushed to:
https://hub.docker.com/r/statusteam/p2p-health-bot/

# License

MIT
