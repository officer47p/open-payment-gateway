package evm

type Network struct {
	Name                string `json:"name"`
	Currency            string `json:"currency"`
	ChainID             int64  `json:"chainID"`
	Decimals            int64  `json:"decimals"`
	StartingBlockNumber int64  `json:"startingBlockNumber"`
}

type NetworkConfig struct {
	Network   Network    `json:"network"`
	Contracts []Contract `json:"contracts"`
}

type Contract struct {
	Name                string `json:"name"`
	Currency            string `json:"currency"`
	Decimals            int    `json:"decimals"`
	ContractAddress     string `json:"contractAddress"`
	Standard            string `json:"standard"`
	StartingBlockNumber int64  `json:"startingBlockNumber"`
}
