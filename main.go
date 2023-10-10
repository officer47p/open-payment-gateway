package main

import (
	"log"
	"open-payment-gateway/config"
	"open-payment-gateway/db"
	"open-payment-gateway/internal_notification"
	"open-payment-gateway/listeners"
	"open-payment-gateway/providers"
	"open-payment-gateway/types"
	"open-payment-gateway/utils"
	"sync"
	"time"
)

func main() {

	networkConfig, err := config.LoadNetworkConfig("config/network-config.json")
	if err != nil {
		log.Fatalf("[init] error loading network config file: %s", err.Error())

	}
	log.Println("[init] network config is loaded")

	// Loading environment variables
	env, err := utils.LoadEnvVariableFile(".env")
	if err != nil {
		log.Fatalf("[init] error loading .env file: %s", err.Error())
	}
	log.Println("[init] environment variables are loaded")

	// Database connection
	dbClient, err := db.GetDBClient(db.DBClientSettings{
		DBUrl:             db.CreateDBUrl(env.DBUrl, env.DBPort, env.DBName, env.DBUser, env.DBPassword),
		AutoMigrateModels: []any{&types.Block{}, &types.Transaction{}, &types.Address{}},
	})
	if err != nil {
		log.Fatalf("[init] could not connect to the Postgres database: %s", err.Error())
	}
	log.Print("[init] connected to the database")

	// Provider
	provider, err := providers.NewEvmProvider(env.ProviderUrl, networkConfig.Network)
	if err != nil {
		log.Fatalf("[init] could not connect to the provider: %s", err.Error())
	}
	log.Print("[init] provider Initiated")

	// Database Stores
	addressStore := db.NewAddressStore(dbClient)
	blockStore := db.NewBlockStore(dbClient)
	transactionStore := db.NewTransactionStore(dbClient)
	// Internal Service Communication
	internalNotification, err := internal_notification.NewNatsInternalNotification(env.NatsUrl)
	if err != nil {
		log.Fatalf("[init] could not connect to the nats service: %s", err.Error())
	}
	log.Print("[init] connected to the nats service")

	// Listener control channels
	quitch := make(chan struct{})
	wg := &sync.WaitGroup{}
	defer close(quitch)

	// Creating network transaction listener
	evmListener := listeners.NewEvmListener(
		&listeners.EvmListenerConfig{
			Quitch: quitch,
			Wg:     wg,
			// Listener settings, also config
			Network: networkConfig.Network,
			// Stores
			AddressStore:     addressStore,
			BlockStore:       blockStore,
			TransactionStore: transactionStore,
			// Communication
			Notification: internalNotification,
			// Third Parties
			Provider:        provider,
			WaitForNewBlock: time.Second * 1,
		},
	)

	wg.Add(1)
	// Starting the listener
	go evmListener.Start()

	// time.Sleep(time.Millisecond * 1)
	// evmListener.Stop()
	// fmt.Println("loopExitedGracefully")
	wg.Wait()
}
