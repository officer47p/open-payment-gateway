package providers

import (
	"context"
	"math/big"
	"open-payment-gateway/types"

	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EvmProvider struct {
	client  *ethclient.Client
	network types.Network
}

func NewEvmProvider(url string, network types.Network) (EvmProvider, error) {
	ctx := context.Background()
	client, err := ethclient.DialContext(ctx, url)

	if err != nil {
		return EvmProvider{}, err
	}

	return EvmProvider{client: client, network: network}, nil
}

func (p EvmProvider) GetLatestBlockNumber() (int64, error) {
	ctx := context.Background()
	n, err := p.client.BlockNumber(ctx)
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
	ctx := context.Background()
	block, err := p.client.BlockByNumber(ctx, big.NewInt(n))
	if err != nil {
		return types.Block{}, err
	}

	var transactions []types.Transaction
	for _, t := range block.Transactions() {
		transactions = append(transactions, parseTransaction(t, block, p.network))
	}

	return types.Block{
		Network:           p.network.Name,
		BlockNumber:       block.Number().Int64(),
		BlockHash:         block.Hash().String(),
		PreviousBlockHash: block.ParentHash().String(),
		Transactions:      transactions,
	}, nil
}

func parseTransaction(tx *ethereumTypes.Transaction, block *ethereumTypes.Block, network types.Network) types.Transaction {
	blockNumber := block.Number().Int64()
	blockHash := block.Hash().String()
	txHash := tx.Hash().String()
	from, ok := getSenderAddressForTransaction(tx)
	if !ok {
		from = ""
	}
	to, ok := getReceiverAddressForTransaction(tx)
	if !ok {
		to = ""
	}
	value := tx.Value().String()

	return types.Transaction{
		Broadcasted: false,
		Network:     network.Name,
		Currency:    network.Currency,
		BlockNumber: blockNumber,
		BlockHash:   blockHash,
		TxHash:      txHash,
		Value:       value,
		From:        from,
		To:          to,
	}
}

func getSenderAddressForTransaction(t *ethereumTypes.Transaction) (string, bool) {
	address, err := ethereumTypes.Sender(ethereumTypes.LatestSignerForChainID(t.ChainId()), t)
	if err != nil {
		return "", false
	}

	return address.String(), true
}

func getReceiverAddressForTransaction(t *ethereumTypes.Transaction) (string, bool) {
	if t.To() != nil {
		return t.To().String(), true
	}
	return "", false
}
