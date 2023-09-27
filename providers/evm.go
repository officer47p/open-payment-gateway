package providers

import (
	"context"
	"errors"
	"open-payment-gateway/types"

	"github.com/ethereum/go-ethereum/ethclient"
)

type EvmProvider struct {
	client *ethclient.Client
}

func (p EvmProvider) GetLatestBlockNumber() (int64, error) {
	context := context.Background()
	n, err := p.client.BlockNumber(context)
	if err != nil {
		return 0, err
	}
	if n < 0 {
		return 0, errors.New("node is not synced")
	}

	return int64(n), nil
}

func (p EvmProvider) GetBlockByNumber(n int64) (types.Block, error) {
	return types.Block{BlockNumber: n, BlockHash: "dfkjnskfns", PreviousBlockHash: "dfkjndskfndsjn"}, nil
}

func NewEvmProvider(url string) (EvmProvider, error) {
	context := context.Background()
	client, err := ethclient.DialContext(context, url)

	if err != nil {
		return EvmProvider{}, err
	}

	return EvmProvider{client: client}, nil
}
