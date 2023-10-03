package db

import (
	"open-payment-gateway/types"

	"gorm.io/gorm"
)

type TransactionStore interface {
	SaveTransaction(*types.Transaction) error
}

type SQLTransactionStore struct {
	client *gorm.DB
}

func NewTransactionStore(c *gorm.DB) *SQLTransactionStore {
	return &SQLTransactionStore{client: c}
}

func (s SQLTransactionStore) SaveTransaction(t *types.Transaction) error {
	result := s.client.Create(t)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
