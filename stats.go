package main

import (
	"fmt"
	"sync"
	"time"
)

// Stats represents messages' statistics.
type Stats struct {
	mx       sync.RWMutex
	sent     int
	received int
	delays   []time.Duration
}

// NewStats returns new empty Stats object.
func NewStats() *Stats {
	return &Stats{}
}

// AddSent adds information about sent messages.
func (s *Stats) AddSent() {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.sent++
}

// AddRoundtrip adds information about successful message roundtrip.
func (s *Stats) AddRountrip(d time.Duration) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.received++
	s.delays = append(s.delays, d)
}

// Print dumps stats to the console.
func (s *Stats) Print() {
	s.mx.RLock()
	defer s.mx.RUnlock()
	fmt.Println("-------------------------")
	fmt.Println("Time:", time.Now())
	fmt.Println("Sent:", s.sent)
	fmt.Println("Received:", s.received)
	fmt.Println("Delays:", s.delays)
	fmt.Println("-------------------------")
}
