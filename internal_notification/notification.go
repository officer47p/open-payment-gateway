package internal_notification

import (
	"time"

	"github.com/nats-io/nats.go"
)

type InternalNotification interface {
	Notify(string, string) error
	// Close() error
}

type NatsInternalNotification struct {
	client *nats.Conn
}

func NewNatsInternalNotification(url string) (*NatsInternalNotification, error) {
	// TODO: Add close method to the connection
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}

	return &NatsInternalNotification{client: nc}, nil
}

// func (n *NatsInternalNotification) Close() error {
// 	err := n.client.Drain()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (n *NatsInternalNotification) Notify(subject string, v string) error {
	_, err := n.client.Request(subject, []byte(v), time.Second)
	if err != nil {
		return err
	}
	return nil
}
