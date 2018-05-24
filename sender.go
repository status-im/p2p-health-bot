package main

import (
	"log"
	"time"

	"github.com/status-im/status-go-sdk"
)

func startSender(ch *sdk.Channel, interval time.Duration, statsPort string) {
	var (
		counter int
		ticker  = time.NewTicker(interval)
		pending = make(map[int]time.Time)
		recvCh  = make(chan Msg, 1000)
	)

	if _, err := ch.Subscribe(func(m *sdk.Msg) {
		rawmsg, ok := m.Properties.(*sdk.PublishMsg)
		if !ok {
			log.Println("Wrong message props type received: %T", m.Properties)
			return
		}
		msg := Msg(rawmsg.Text)
		if msg.IsResponse() {
			recvCh <- msg
		}
	}); err != nil {
		log.Fatal(err)
	}

	stats := NewStats(statsPort)

	for {
		select {
		case <-ticker.C:
			var body = NewRequestMsg(counter)
			err := ch.Publish(string(body))
			if err != nil {
				log.Printf("[ERROR] Failed to send health message: %s", err)
				continue
			}
			pending[counter] = time.Now()
			counter++
			stats.AddSent()
		case msg := <-recvCh:
			c, err := msg.Counter()
			if err != nil {
				log.Printf("[ERROR] Failed to parse health message: %s", err)
				continue
			}
			start, ok := pending[c]
			if !ok {
				log.Printf("[ERROR] Received response for counter never sent (another sender bot running?): %s", err)
				continue
			}
			delete(pending, c)
			dur := time.Since(start)
			stats.AddRountrip(dur)
		}
	}
}
