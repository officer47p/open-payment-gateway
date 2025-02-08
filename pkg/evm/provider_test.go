package evm

// import (
// 	"context"
// 	"math/big"
// 	"open-payment-gateway/types"
// 	"testing"
// )

// // MockEthClient is a mock implementation of the ethclient.Client interface for testing purposes.
// type MockEthClient struct{}

// func (m *MockEthClient) BlockNumber(ctx context.Context) (uint64, error) {
// 	return 1234, nil
// }

// func (m *MockEthClient) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
// 	return &types.Block{
// 		BlockNumber:       1234,
// 		BlockHash:         "blockHash",
// 		PreviousBlockHash: "parentHash",
// 		Transactions:      []types.Transaction{},
// 	}, nil
// }

// func TestEvmProvider_GetLatestBlockNumber(t *testing.T) {
// 	// Create a new EvmProvider with a mock EthClient
// 	provider := EvmProvider{
// 		client:  &MockEthClient{},
// 		network: types.Network{},
// 	}

// 	blockNumber, err := provider.GetLatestBlockNumber()
// 	if err != nil {
// 		t.Errorf("GetLatestBlockNumber failed with error: %v", err)
// 	}

// 	if blockNumber != 1234 {
// 		t.Errorf("Expected block number 1234, but got %v", blockNumber)
// 	}
// }

// func TestEvmProvider_GetBlockByNumber(t *testing.T) {
// 	// Create a new EvmProvider with a mock EthClient
// 	provider := EvmProvider{
// 		client:  &MockEthClient{},
// 		network: types.Network{},
// 	}

// 	blockNumber := int64(1234)

// 	block, err := provider.GetBlockByNumber(blockNumber)
// 	if err != nil {
// 		t.Errorf("GetBlockByNumber failed with error: %v", err)
// 	}

//		if block.Network != "" || block.BlockNumber != 1234 {
//			t.Errorf("Expected block fields to be set, but got %+v", block)
//		}
//	}
// package providers

// import (
// 	"math/big"
// 	"open-payment-gateway/types"
// 	"testing"

// 	"github.com/ethereum/go-ethereum/core/types"
// 	"go.uber.org/mock/gomock"
// )

// func TestEvmProvider_GetLatestBlockNumber(t *testing.T) {
// 	// Create a new Gomock controller
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	// Create a mock ethclient.Client
// 	mockEthClient := mocks.NewMockClient(ctrl)

// 	// Define the expected behavior of the mock
// 	mockEthClient.EXPECT().BlockNumber(gomock.Any()).Return(uint64(1234), nil)

// 	// Create a new EvmProvider with the mock EthClient
// 	provider := EvmProvider{
// 		client:  mockEthClient,
// 		network: types.Network{},
// 	}

// 	blockNumber, err := provider.GetLatestBlockNumber()
// 	if err != nil {
// 		t.Errorf("GetLatestBlockNumber failed with error: %v", err)
// 	}

// 	if blockNumber != 1234 {
// 		t.Errorf("Expected block number 1234, but got %v", blockNumber)
// 	}
// }

// func TestEvmProvider_GetBlockByNumber(t *testing.T) {
// 	// Create a new Gomock controller
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	// Create a mock ethclient.Client
// 	mockEthClient := mocks.NewMockClient(ctrl)

// 	// Define the expected behavior of the mock
// 	mockEthClient.EXPECT().BlockByNumber(gomock.Any(), gomock.Any()).Return(&types.Block{
// 		Number:     big.NewInt(1234),
// 		Hash:       types.NewHash([]byte("blockHash")),
// 		ParentHash: types.NewHash([]byte("parentHash")),
// 	}, nil)

// 	// Create a new EvmProvider with the mock EthClient
// 	provider := EvmProvider{
// 		client:  mockEthClient,
// 		network: types.Network{},
// 	}

// 	blockNumber := int64(1234)

// 	block, err := provider.GetBlockByNumber(blockNumber)
// 	if err != nil {
// 		t.Errorf("GetBlockByNumber failed with error: %v", err)
// 	}

// 	if block.Network != "" || block.BlockNumber != 1234 {
// 		t.Errorf("Expected block fields to be set, but got %+v", block)
// 	}
// }
