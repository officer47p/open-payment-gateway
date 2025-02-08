package config

import (
	"encoding/json"
	"fmt"
	"io"
	"open-payment-gateway/pkg/evm"
	"os"
)

func LoadNetworkConfig(path string) (evm.NetworkConfig, error) {
	fileContent, err := os.Open(path)

	if err != nil {
		return evm.NetworkConfig{}, err
	}

	defer fileContent.Close()

	byteResult, err := io.ReadAll(fileContent)

	if err != nil {
		return evm.NetworkConfig{}, err
	}

	var networkConfig evm.NetworkConfig

	json.Unmarshal(byteResult, &networkConfig)
	fmt.Printf("network: %+v,\ncontracts: %+v\n", networkConfig.Network, networkConfig.Contracts)

	return networkConfig, nil
}
