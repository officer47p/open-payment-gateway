package db

import (
	"open-payment-gateway/types"
)

type DB interface {
	AddressExists(a types.Address) (bool, error)
	GetLatestProcessedBlockNumber() (int64, error)
	SaveBlock(b types.Block) error
}

type PostgresDB struct{}

func (db PostgresDB) AddressExists(a types.Address) (bool, error) {
	return true, nil
}

func (db PostgresDB) GetLatestProcessedBlockNumber() (int64, error) {
	return 432432, nil
}

func (db PostgresDB) SaveBlock(b types.Block) error {
	return nil
}
