package main

import (
	"log"
	"open-payment-gateway/db"
	"open-payment-gateway/listeners"
	"open-payment-gateway/providers"
	"open-payment-gateway/types"
	"open-payment-gateway/utils"
	"os"
	"strconv"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	var network types.Network = types.Network{}

	err := utils.LoadEnvVariableFile(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	dbURL := db.CreateDBUrl(os.Getenv("DB_URL"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))
	postgresClient, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the Postgres database")
	}
	postgresClient.AutoMigrate(&types.Block{}, &types.Transaction{}, &types.Address{})

	provider, err := providers.NewEvmProvider(os.Getenv("PROVIDER_URL"))
	if err != nil {
		panic("Could not connect to the provider")

	}

	chainId, err := strconv.ParseInt(os.Getenv("CHAIN_ID"), 10, 64)
	if err != nil {
		panic("Could not parse chain id env variables")
	}

	decimals, err := strconv.ParseInt(os.Getenv("DECIMALS"), 10, 64)
	if err != nil {
		panic("Could not parse decimals env variables")
	}

	network.Name = os.Getenv("NETWORK_NAME")
	network.Currency = os.Getenv("NETWORK_CURRENCY")
	network.ChainID = chainId
	network.Decimals = decimals

	startingBlockNumber, err := strconv.ParseInt(os.Getenv("STARTING_BLOCK_NUMBER"), 10, 64)
	if err != nil {
		panic("Could not parse the starting block number")
	}

	addressStore := db.NewAddressStore(postgresClient)
	blockStore := db.NewBlockStore(postgresClient)
	transactionStore := db.NewTransactionStore(postgresClient)
	quitch := make(chan struct{})
	wg := &sync.WaitGroup{}
	defer close(quitch)

	evmListener := listeners.NewEvmListener(
		quitch,
		wg,
		network,
		startingBlockNumber,
		addressStore,
		blockStore,
		transactionStore,
		provider,
	)

	wg.Add(1)
	go evmListener.Start()

	// time.Sleep(time.Millisecond * 1)
	// evmListener.Stop()
	// fmt.Println("loopExitedGracefully")
	wg.Wait()
}
