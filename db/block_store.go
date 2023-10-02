package db

import (
	"open-payment-gateway/types"

	"gorm.io/gorm"
)

type BlockStore interface {
	SaveBlock(b *types.Block) error
	GetLatestProcessedBlockNumber() (int64, error)
}

type SQLBlockStore struct {
	client *gorm.DB
}

func NewBlockStore(c *gorm.DB) *SQLBlockStore {
	return &SQLBlockStore{client: c}
}

func (s *SQLBlockStore) SaveBlock(b *types.Block) error {
	result := s.client.Create(b)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *SQLBlockStore) GetLatestProcessedBlockNumber() (int64, error) {
	var latestProcessedBlockNumber int64
	result := s.client.Model(&types.Block{}).Select("max(block_number)")

	if result.Error != nil {
		return 0, result.Error
	}

	err := result.Row().Scan(&latestProcessedBlockNumber)
	if err != nil {
		return -1, nil
	}

	return latestProcessedBlockNumber, nil
}
