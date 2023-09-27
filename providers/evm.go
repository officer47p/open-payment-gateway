package providers

import "open-payment-gateway/types"

type EvmProvider struct {
	PrividerUrl string
}

func (p EvmProvider) GetLatestBlockNumber() (int64, error) {
	return 100, nil
}

func (p EvmProvider) GetBlockByNumber(int64) (types.Block, error) {
	return types.Block{}, nil
}

func NewEvmProvider(url string) EvmProvider {
	return EvmProvider{PrividerUrl: url}
}
