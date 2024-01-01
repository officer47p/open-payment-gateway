package listeners

import (
	"open-payment-gateway/stores"
	"open-payment-gateway/types"
	"strings"
	"sync"
	"testing"
	"time"

	"golang.org/x/exp/slices"
)

// Mocking AddressStore
type MockAddressStore struct {
	addressExistsCalls int64
	addresses          []string
}

func (s *MockAddressStore) AddressExists(address string) (bool, error) {
	s.addressExistsCalls++
	exists := slices.Contains(s.addresses, strings.ToLower(address))
	return exists, nil
}

// Mock BlockStore
type MockBlockStore struct {
	saveBlockCalls                     int64
	blocks                             []types.Block
	getLatestProcessedBlockNumberCalls int64
}

func (s *MockBlockStore) SaveBlock(b *types.Block) error {
	s.saveBlockCalls++
	s.blocks = append(s.blocks, *b)
	return nil
}

func (s *MockBlockStore) GetLatestProcessedBlockNumber() (int64, error) {
	s.getLatestProcessedBlockNumberCalls++
	if len(s.blocks) == 0 {
		return -1, nil
	}
	latestProcessedBlockNumber := 0
	for _, b := range s.blocks {
		if b.BlockNumber > int64(latestProcessedBlockNumber) {
			latestProcessedBlockNumber = int(b.BlockNumber)
		}
	}
	return int64(latestProcessedBlockNumber), nil
}

type MockTransactionStore struct {
}

func (s MockTransactionStore) SaveTransaction(t *stores.Transaction) error {
	return nil
}

func (s MockTransactionStore) UpdateBroadcasted(txHash string, value bool) error {
	return nil
}

var network = types.Network{
	Name:     "ethereum",
	Currency: "ETH",
	ChainID:  1,
	Decimals: 18,
}

type MockInternalNotification struct {
}

func (in MockInternalNotification) Notify(string, string) error {
	return nil
}

type MockEvmProvider struct {
}

func (p MockEvmProvider) GetLatestBlockNumber() (int64, error) {
	return int64(33), nil
}

func (p MockEvmProvider) GetBlockByNumber(n int64) (types.Block, error) {
	return types.Block{}, nil
}

func TestEvmListenerStopFunction(t *testing.T) {
	quitch := make(chan struct{})
	wg := sync.WaitGroup{}
	mockAddressStore := MockAddressStore{}
	mockBlockStore := MockBlockStore{}
	mockTransactionStore := MockTransactionStore{}
	mockNotification := MockInternalNotification{}
	mockProvider := MockEvmProvider{}

	l := NewEvmListener(&EvmListenerConfig{
		Quitch:           quitch,
		Wg:               &wg,
		Network:          network,
		AddressStore:     &mockAddressStore,
		BlockStore:       &mockBlockStore,
		TransactionStore: &mockTransactionStore,
		Notification:     &mockNotification,
		Provider:         mockProvider,
		WaitForNewBlock:  time.Second,
	})

	wg.Add(1)
	go l.Start()

	stopped := l.Stop()

	wg.Wait()
	if !stopped {
		t.Error("expected the listener to stop, but returned false")
	}

	if calls := mockAddressStore.addressExistsCalls; calls != 0 {
		t.Errorf("expected zero calls to address store, got: %d", calls)
	}

	if calls := mockBlockStore.getLatestProcessedBlockNumberCalls; calls != 0 {
		t.Errorf("expected zero calls to block store, got: %d", calls)
	}

}
