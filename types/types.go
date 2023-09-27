package types

import "gorm.io/gorm"

type Address string

type Network string

type Currency string

type Block struct {
	gorm.Model
	BlockNumber       int64
	BlockHash         string
	PreviousBlockHash string
	Transactions      []Transaction `gorm:"-"`
}

type Transaction struct {
	gorm.Model
	Txhash string
	Value  int64
	From   Address
	To     Address
}
