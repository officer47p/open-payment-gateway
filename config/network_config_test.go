package config

import (
	"io/ioutil"
	"open-payment-gateway/types"
	"os"
	"testing"
)

func TestLoadNetworkConfig(t *testing.T) {
	// Create a temporary JSON configuration file for testing
	tempFile := createTempConfigFile(t)
	defer os.Remove(tempFile.Name())

	// Load the network configuration from the temporary file
	networkConfig, err := LoadNetworkConfig(tempFile.Name())
	if err != nil {
		t.Errorf("Error loading network configuration: %v", err)
	}

	// Verify the loaded network configuration
	expectedNetwork := types.Network{
		Name:                "ethereum",
		Currency:            "ETH",
		ChainID:             1,
		Decimals:            18,
		StartingBlockNumber: 9797391,
	}

	if networkConfig.Network != expectedNetwork {
		t.Errorf("Loaded network configuration is not as expected, got %+v, want %+v", networkConfig.Network, expectedNetwork)
	}

	expectedContracts := []types.Contract{
		{
			Name:                "Tether",
			Currency:            "USDT",
			Decimals:            18,
			ContractAddress:     "0x4723956743657482936587326",
			Standard:            "ERC20",
			StartingBlockNumber: 9797380,
		},
	}

	if len(networkConfig.Contracts) != len(expectedContracts) {
		t.Errorf("Number of contracts in the loaded configuration does not match")
	}

	// Additional checks can be added to compare the contracts as well
}

func createTempConfigFile(t *testing.T) *os.File {
	// Create a temporary JSON configuration file for testing
	configData := `{
		"network": {
			"name": "ethereum",
			"currency": "ETH",
			"chainID": 1,
			"decimals": 18,
			"startingBlockNumber": 9797391
		},
		"contracts": [
			{
				"name": "Tether",
				"currency": "USDT",
				"decimals": 18,
				"contractAddress": "0x4723956743657482936587326",
				"standard": "ERC20",
				"startingBlockNumber": 9797380
			}
		]
	}`

	tempFile, err := ioutil.TempFile("", "test_config_*.json")
	if err != nil {
		t.Fatalf("Failed to create temporary config file: %v", err)
	}

	if _, err := tempFile.WriteString(configData); err != nil {
		t.Fatalf("Failed to write to temporary config file: %v", err)
	}

	return tempFile
}
