package main

import (
	"log"
	"open-payment-gateway/db"
	"open-payment-gateway/internal_notification"
	"open-payment-gateway/listeners"
	"open-payment-gateway/providers"
	"open-payment-gateway/types"
	"open-payment-gateway/utils"
	"sync"
)

func main() {
	// Loading environment variables
	env, err := utils.LoadEnvVariableFile(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Creating network object
	network := types.Network{
		Name:     env.NetworkName,
		Currency: env.NetworkCurrency,
		ChainID:  env.ChainID,
		Decimals: env.Decimals,
	}

	// Database connection
	dbClient, err := db.GetDBClient(db.DBClientSettings{
		DBUrl:             db.CreateDBUrl(env.DBUrl, env.DBPort, env.DBName, env.DBUser, env.DBPassword),
		AutoMigrateModels: []any{&types.Block{}, &types.Transaction{}, &types.Address{}},
	})
	if err != nil {
		panic("Could not connect to the Postgres database")
	}

	// Provider
	provider, err := providers.NewEvmProvider(env.ProviderUrl)
	if err != nil {
		panic("Could not connect to the provider")
	}

	// Database Stores
	addressStore := db.NewAddressStore(dbClient)
	blockStore := db.NewBlockStore(dbClient)
	transactionStore := db.NewTransactionStore(dbClient)

	// Internal Service Communication
	internalNotification := internal_notification.NewNatsInternalNotification(transactionStore)

	// Listener control channels
	quitch := make(chan struct{})
	wg := &sync.WaitGroup{}
	defer close(quitch)

	// Creating network transaction listener
	evmListener := listeners.NewEvmListener(
		quitch,
		wg,
		network,
		env.StartingBlockNumber,
		addressStore,
		blockStore,
		transactionStore,
		internalNotification,
		provider,
	)

	wg.Add(1)
	// Starting the listener
	go evmListener.Start()

	// time.Sleep(time.Millisecond * 1)
	// evmListener.Stop()
	// fmt.Println("loopExitedGracefully")
	wg.Wait()
}
