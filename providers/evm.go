package providers

import (
	"context"
	"math/big"
	"open-payment-gateway/types"

	ethereum_types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EvmProvider struct {
	client *ethclient.Client
}

func NewEvmProvider(url string) (EvmProvider, error) {
	context := context.Background()
	client, err := ethclient.DialContext(context, url)

	if err != nil {
		return EvmProvider{}, err
	}

	return EvmProvider{client: client}, nil
}

func (p EvmProvider) GetLatestBlockNumber() (int64, error) {
	context := context.Background()
	n, err := p.client.BlockNumber(context)
	if err != nil {
		return 0, err
	}

	// Return type of BlockNumber() is uint64, so it will never be a negative number
	// if n < 0 {
	// 	return 0, errors.New("node is not synced")
	// }

	return int64(n), nil
}

func (p EvmProvider) GetBlockByNumber(n int64) (types.Block, error) {
	context := context.Background()
	block, err := p.client.BlockByNumber(context, big.NewInt(n))
	if err != nil {
		return types.Block{}, err
	}

	transactions := []types.Transaction{}
	for _, t := range block.Transactions() {
		transactions = append(transactions, parseTransaction(t, block))
	}

	return types.Block{BlockNumber: block.Number().Int64(), BlockHash: block.Hash().String(), PreviousBlockHash: block.ParentHash().String(), Transactions: transactions}, nil
}

func parseTransaction(tx *ethereum_types.Transaction, block *ethereum_types.Block) types.Transaction {
	blockNumber := block.Number().Int64()
	blockHash := block.Hash().String()
	txhash := tx.Hash().String()
	from, ok := getSenderAddressForTransaction(tx)
	if !ok {
		from = ""
	}
	to, ok := getReceiverAddressForTransaction(tx)
	if !ok {
		to = ""
	}
	value := tx.Value().String()

	return types.Transaction{BlockNumber: blockNumber, BlockHash: blockHash, Txhash: txhash, Value: value, From: from, To: to}
}

func getSenderAddressForTransaction(t *ethereum_types.Transaction) (string, bool) {
	address, err := ethereum_types.Sender(ethereum_types.LatestSignerForChainID(t.ChainId()), t)
	if err != nil {
		return "", false
	}

	return address.String(), true
}

func getReceiverAddressForTransaction(t *ethereum_types.Transaction) (string, bool) {
	if t.To() != nil {
		return t.To().String(), true
	}
	return "", false
}
