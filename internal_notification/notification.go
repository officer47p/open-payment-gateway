package internal_notification

import (
	"fmt"
	"open-payment-gateway/db"
)

type InternalNotification interface {
	Notify(string, string) error
}

type NatsInternalNotification struct {
	transactionStore db.TransactionStore
}

func NewNatsInternalNotification(s db.TransactionStore) *NatsInternalNotification {
	return &NatsInternalNotification{
		transactionStore: s,
	}
}

func (n *NatsInternalNotification) Notify(topic string, v string) error {
	fmt.Printf("Sent Internal Notification. Topic: %s, Value: %s", topic, v)
	return nil
}
