package internal_notification

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

type InternalNotification interface {
	Notify(string, string) error
}

type NatsInternalNotification struct {
	client *nats.Conn
}

func NewNatsInternalNotification(url string) (*NatsInternalNotification, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}

	return &NatsInternalNotification{client: nc}, nil
}

func (n *NatsInternalNotification) Notify(subject string, v string) error {
	if err := n.client.Publish(subject, []byte(v)); err != nil {
		return err
	}
	fmt.Printf("Sent Internal Notification. Topic: %s, Value: %s", subject, v)
	return nil
}
