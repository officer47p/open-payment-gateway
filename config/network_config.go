package config

import (
	"encoding/json"
	"fmt"
	"io"
	"open-payment-gateway/types"
	"os"
)

func LoadNetworkConfig(path string) (types.NetworkConfig, error) {
	fileContent, err := os.Open(path)

	if err != nil {
		return types.NetworkConfig{}, err
	}

	defer fileContent.Close()

	byteResult, err := io.ReadAll(fileContent)

	if err != nil {
		return types.NetworkConfig{}, err
	}

	var networkConfig types.NetworkConfig

	json.Unmarshal(byteResult, &networkConfig)
	fmt.Printf("network: %+v,\ncontracts: %+v\n", networkConfig.Network, networkConfig.Contracts)

	return networkConfig, nil
}
