package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Msg represents health check message.
type Msg string

// NewRequestMsg constructs new request health message.
func NewRequestMsg(counter int) Msg {
	return Msg(fmt.Sprintf("Health Check Request|%d", counter))
}

// NewResponseMsg constructs new request health message.
func NewResponseMsg(counter int) Msg {
	return Msg(fmt.Sprintf("Health Check Response|%d", counter))
}

// Counter returns counter value from health message.
func (m *Msg) Counter() (int, error) {
	fields := strings.Split(string(*m), "|")
	if len(fields) != 2 {
		return 0, fmt.Errorf("wrong length: %s", *m)
	}
	c, err := strconv.ParseInt(fields[1], 10, 0)
	if err != nil {
		return 0, fmt.Errorf("wrong counter value: %s", err)
	}
	return int(c), nil
}

func (m *Msg) IsRequest() bool {
	return strings.Contains(string(*m), "Health Check Request")
}

func (m *Msg) IsResponse() bool {
	return strings.Contains(string(*m), "Health Check Response")
}
