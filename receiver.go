package main

import (
	"log"

	"github.com/status-im/status-go-sdk"
)

func startReceiver(ch *sdk.Channel) {
	if _, err := ch.Subscribe(func(m *sdk.Msg) {
		log.Println("[DEBUG] Got message: %v", m)
		rawmsg, ok := m.Properties.(*sdk.PublishMsg)
		if !ok {
			log.Println("[ERROR] Wrong message props type received: %T", m.Properties)
			return
		}
		msg := Msg(rawmsg.Text)
		if msg.IsRequest() {
			counter, err := msg.Counter()
			if err != nil {
				log.Println("[ERROR] Can't extract counter: %v", err)
			}
			go func(counter int) {
				var body = NewResponseMsg(counter)
				if err := ch.Publish(string(body)); err != nil {
					log.Println("[ERROR] Can't send response: %v", err)
				}
			}(counter)
		}
	}); err != nil {
		log.Fatal(err)
	}
}
