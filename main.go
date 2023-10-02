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
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Find .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	dbURL := db.CreateDBUrl(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_URL"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	postgresClient, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the Postgres database")
	}
	postgresClient.AutoMigrate(&types.Block{}, &types.Transaction{})

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

	addressStore := db.NewAddressStore(postgresClient)
	blockStore := db.NewBlockStore(postgresClient)

	evmListener := listeners.NewEvmListener(
		os.Getenv("NETWORK_NAME"),
		os.Getenv("NETWORK_CURRENCY"),
		chainId,
		startingBlockNumber,
		addressStore,
		blockStore,
		provider,
	)

	evmListener.Start()
}
