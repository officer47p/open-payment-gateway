package evm

import (
	"encoding/json"
	"testing"
)

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
