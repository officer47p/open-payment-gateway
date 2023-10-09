package types

import (
	"gorm.io/gorm"
)

type Block struct {
	gorm.Model
	ID                uint          `gorm:"column:id;primaryKey;autoIncrement"`
	Network           string        `gorm:"column:network;not null"`
	BlockNumber       int64         `gorm:"column:block_number;not null;unique"`
	BlockHash         string        `gorm:"column:block_hash;not null;unique"`
	PreviousBlockHash string        `gorm:"column:previous_block_hash;not null;unique"`
	Transactions      []Transaction `gorm:"-"`
}

type Transaction struct {
	gorm.Model
	ID          uint   `gorm:"column:id;primaryKey;autoIncrement"`
	Broadcasted bool   `gorm:"column:broadcasted;not null"`
	BlockNumber int64  `gorm:"column:block_number;not null"`
	BlockHash   string `gorm:"column:block_hash;not null"`
	Network     string `gorm:"column:network;not null"`
	Currency    string `gorm:"column:currency;not null"`
	TxHash      string `gorm:"column:tx_hash;not null;unique"`
	TxType      string `gorm:"column:tx_type;not null"`
	Value       string `gorm:"column:value;not null"`
	From        string `gorm:"column:from;not null"`
	To          string `gorm:"column:to;not null"`
}

type Address struct {
	gorm.Model
	ID      uint   `gorm:"column:id;primaryKey;autoIncrement"`
	Address string `gorm:"column:address;not null;unique"`
	HDPath  string `gorm:"column:hd_path;not null;unique"`
}

type Network struct {
	Name     string
	Currency string
	ChainID  int64
	Decimals int64
}
