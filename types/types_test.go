package types

import (
	"encoding/json"
	"testing"
)

func TestNewTransactionNotificationToJSON(t *testing.T) {
	notification := NewTransactionNotification{
		BlockNumber: 12345,
		BlockHash:   "0x123abc",
		Network:     "Mainnet",
		Currency:    "ETH",
		TxHash:      "0x456def",
		TxType:      "Transfer",
		Value:       "10.5",
		From:        "0xabcdef123",
		To:          "0x789ghi456",
	}

	expectedJSON := `{"block_number":12345,"block_hash":"0x123abc","network":"Mainnet","currency":"ETH","tx_hash":"0x456def","tx_type":"Transfer","value":"10.5","from":"0xabcdef123","to":"0x789ghi456"}`

	jsonStr, err := notification.ToJSON()
	if err != nil {
		t.Errorf("Error converting NewTransactionNotification to JSON: %v", err)
	}

	if jsonStr != expectedJSON {
		t.Errorf("Expected JSON: %s\nActual JSON: %s", expectedJSON, jsonStr)
	}
}
func TestNetworkConfigJSONRoundTrip(t *testing.T) {
	networkConfig := NetworkConfig{
		Network: Network{
			Name:                "MyNetwork",
			Currency:            "MYC",
			ChainID:             123,
			Decimals:            18,
			StartingBlockNumber: 1000,
		},
		Contracts: []Contract{
			{
				Name:                "MyToken",
				Currency:            "MYT",
				Decimals:            18,
				ContractAddress:     "0x123abc",
				Standard:            "ERC20",
				StartingBlockNumber: 1000,
			},
		},
	}

	// Convert NetworkConfig to JSON
	jsonData, err := json.Marshal(networkConfig)
	if err != nil {
		t.Errorf("Error marshaling NetworkConfig to JSON: %v", err)
	}

	// Convert JSON back to NetworkConfig
	var parsedConfig NetworkConfig
	err = json.Unmarshal(jsonData, &parsedConfig)
	if err != nil {
		t.Errorf("Error unmarshaling JSON to NetworkConfig: %v", err)
	}

	// Check if the original and parsed NetworkConfig are equal
	if !compareNetworkConfigs(networkConfig, parsedConfig) {
		t.Errorf("Original NetworkConfig and parsed NetworkConfig do not match.")
	}
}

func compareNetworkConfigs(a, b NetworkConfig) bool {
	if a.Network != b.Network {
		return false
	}
	if len(a.Contracts) != len(b.Contracts) {
		return false
	}
	for i := range a.Contracts {
		if a.Contracts[i] != b.Contracts[i] {
			return false
		}
	}
	return true
}
