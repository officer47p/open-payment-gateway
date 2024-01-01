package providers

import (
	"context"
	"math/big"
	"open-payment-gateway/stores"
	"open-payment-gateway/types"

	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EvmProvider interface {
	GetLatestBlockNumber() (int64, error)
	GetBlockByNumber(n int64) (types.Block, error)
}

type ThirdPartyEvmProvider struct {
	client  *ethclient.Client
	network types.Network
}

func NewEvmProvider(url string, network types.Network) (ThirdPartyEvmProvider, error) {
	ctx := context.Background()
	client, err := ethclient.DialContext(ctx, url)

	if err != nil {
		return ThirdPartyEvmProvider{}, err
	}

	return ThirdPartyEvmProvider{client: client, network: network}, nil
}

func (p ThirdPartyEvmProvider) GetLatestBlockNumber() (int64, error) {
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

func (p ThirdPartyEvmProvider) GetBlockByNumber(n int64) (types.Block, error) {
	ctx := context.Background()
	block, err := p.client.BlockByNumber(ctx, big.NewInt(n))
	if err != nil {
		return types.Block{}, err
	}

	var transactions []stores.Transaction
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

func parseTransaction(tx *ethereumTypes.Transaction, block *ethereumTypes.Block, network types.Network) stores.Transaction {
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

	return stores.Transaction{
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
