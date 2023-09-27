package main

import (
	"log"
	"open-payment-gateway/db"
	"open-payment-gateway/listeners"
	"open-payment-gateway/providers"
	"open-payment-gateway/types"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	// Find .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	database, err := db.NewSQLiteDBClient(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic("Could not connect to the SQLite database")
	}

	provider, err := providers.NewEvmProvider(os.Getenv("PROVIDER_URL"))
	if err != nil {
		panic("Could not connect to the provider")

	}

	chainId, err := strconv.ParseInt(os.Getenv("CHAIN_ID"), 10, 64)
	if err != nil {
		panic("Could not parse the chain id")
	}

	startingBlockNumber, err := strconv.ParseInt(os.Getenv("STARTING_BLOCK_NUMBER"), 10, 64)
	if err != nil {
		panic("Could not parse the starting block number")
	}

	evmListener := listeners.NewEvmListener(
		types.Network(os.Getenv("NETWORK_NAME")),
		types.Currency(os.Getenv("NETWORK_CURRENCY")),
		chainId,
		startingBlockNumber,
		database,
		provider,
	)

	evmListener.Start()
}
