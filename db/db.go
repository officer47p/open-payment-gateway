package db

import (
	"open-payment-gateway/types"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB interface {
	AddressExists(a string) (bool, error)
	GetLatestProcessedBlockNumber() (int64, error)
	SaveBlock(b *types.Block) error
}

type SQLiteDB struct {
	client *gorm.DB
}

func NewSQLiteDBClient(p string) (*SQLiteDB, error) {
	db, err := gorm.Open(sqlite.Open(p), &gorm.Config{})
	if err != nil {
		return &SQLiteDB{}, err
	}

	// Migrate the schema
	db.AutoMigrate(&types.Block{}, &types.Transaction{})

	return &SQLiteDB{client: db}, nil
}

func (db *SQLiteDB) AddressExists(a string) (bool, error) {
	if strings.EqualFold(a, "0x6eb94F4C9CeDF3637f9F3ec21e91231fB8482278") {
		return true, nil
	}
	return false, nil
}

func (db *SQLiteDB) GetLatestProcessedBlockNumber() (int64, error) {
	var latestProcessedBlockNumber int64
	result := db.client.Model(&types.Block{}).Select("max(block_number)")

	if result.Error != nil {
		return 0, result.Error
	}

	err := result.Row().Scan(&latestProcessedBlockNumber)
	if err != nil {
		return -1, nil
	}

	return latestProcessedBlockNumber, nil
}

func (db *SQLiteDB) SaveBlock(b *types.Block) error {
	result := db.client.Create(b)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
