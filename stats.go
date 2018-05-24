package main

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	msgsSent = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "msgs_sent",
		Help: "Messages sent by bot",
	})
	msgsReceived = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "msgs_received",
		Help: "Message responses received by bot",
	})
	msgsLatencies = prometheus.NewSummary(prometheus.SummaryOpts{
		Name:       "msgs_responses_latency",
		Help:       "Latencies of responses to bot messages",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	})
)

func init() {
	// Metrics have to be registered to be exposed
	prometheus.MustRegister(msgsSent)
	prometheus.MustRegister(msgsReceived)
	prometheus.MustRegister(msgsLatencies)
}

// Stats represents messages' statistics.
type Stats struct {
}

// NewStats returns new empty Stats object.
func NewStats(statsPort string) *Stats {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe(statsPort, nil))
	}()
	return &Stats{}
}

// AddSent adds information about sent messages.
func (s *Stats) AddSent() {
	msgsSent.Inc()
}

// AddRoundtrip adds information about successful message roundtrip.
func (s *Stats) AddRountrip(d time.Duration) {
	msgsReceived.Inc()
	msgsLatencies.Observe(float64(d / time.Millisecond))
}
