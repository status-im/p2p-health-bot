package main

import (
	"fmt"
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
	msgsLatencies = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "msgs_responses_latency",
		Help:    "Latencies of responses to bot messages in ms",
		Buckets: []float64{100, 500, 1000, 2000, 3000, 4000, 5000, 7500, 10000, 20000, 30000},
	}, []string{"bot"})
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

// For verifying the service is up
func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

// NewStats returns new empty Stats object.
func NewStats(statsPort string) *Stats {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		// Add most trivial healthcheck
		http.HandleFunc("/health", healthHandler)
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
	msgsLatencies.WithLabelValues("default").Observe(float64(d / time.Millisecond))
}
