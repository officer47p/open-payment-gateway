package types

import (
	"gorm.io/gorm"
)

type Block struct {
	gorm.Model
	BlockNumber       int64
	BlockHash         string
	PreviousBlockHash string
	Transactions      []Transaction `gorm:"-"`
}

type Transaction struct {
	gorm.Model
	BlockNumber int64
	BlockHash   string
	Txhash      string
	TxType      string
	Value       string
	From        string
	To          string
}

type Address struct {
	gorm.Model
	Address string
	HDpath  string
	Balance string
}
