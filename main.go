package main

import (
	"flag"
	"log"

	"time"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/status-im/status-go-sdk"
)

func main() {
	var (
		name      = flag.String("name", "randomstring", "Public chat name used for this health bots")
		interval  = flag.Duration("interval", 5*time.Second, "Interval for health check")
		rpcHost   = flag.String("rpc", "http://localhost:8545", "Host:port to statusd's RPC endpoint")
		statsPort = flag.String("statsPort", ":8080", "Host:port to bind to for exposed Prometheus metrics")
		isSender  = flag.Bool("send", true, "Select bot role, sender or responder")
	)
	flag.Parse()

	rpcClient, err := rpc.Dial(*rpcHost)
	if err != nil {
		log.Fatal(err)
	}

	client := sdk.New(rpcClient)

	a, err := client.SignupAndLogin("password")
	if err != nil {
		log.Fatal(err)
	}

	ch, err := a.JoinPublicChannel(*name)
	if err != nil {
		log.Fatal(err)
	}

	if *isSender {
		startSender(ch, *interval, *statsPort)
	} else {
		startReceiver(ch)
		select {}
	}
}
