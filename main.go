package main

import (
	"open-payment-gateway/db"
	"open-payment-gateway/listeners"
	"open-payment-gateway/providers"
)

func main() {
	evmListener := listeners.NewEvmListener(
		"ethereum",
		"ETH",
		1,
		1000,
		db.PostgresDB{},
		providers.NewEvmProvider("some_url"),
	)

	evmListener.Start()
}
