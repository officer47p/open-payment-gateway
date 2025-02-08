package evm

import (
	"gorm.io/gorm"
)

type TransactionStore interface {
	SaveTransaction(*Transaction) error
	UpdateBroadcasted(txHash string, value bool) error
}

type SQLTransactionStore struct {
	client *gorm.DB
}

func NewTransactionStore(c *gorm.DB) *SQLTransactionStore {
	return &SQLTransactionStore{client: c}
}

func (s SQLTransactionStore) SaveTransaction(t *Transaction) error {
	result := s.client.Create(t)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s SQLTransactionStore) UpdateBroadcasted(txHash string, value bool) error {
	var transaction Transaction
	// TODO(test): Having a case to test case-insesivity of transaction hash and address
	result := s.client.Where("lower(tx_hash) = lower(?)", txHash).First(&transaction)

	if result.Error != nil {
		return result.Error // Transaction not found
	}

	// Update the Broadcasted field
	transaction.Broadcasted = value
	result = s.client.Save(&transaction)

	return result.Error
}
