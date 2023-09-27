package types

type Address string

type Network string

type Currency string

type Block struct {
	BlockNumber       int64
	BlockHash         string
	PreviousBlockHash string
	Transactions      []Transaction
}

type Transaction struct {
	Txhash string
	Value  int64
	From   Address
	To     Address
}
