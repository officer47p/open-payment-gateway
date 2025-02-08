package main

import (
	"log"
	"open-payment-gateway/pkg/config"
	"open-payment-gateway/pkg/db"
	"open-payment-gateway/pkg/eventbus"
	"open-payment-gateway/pkg/evm"

	"sync"
	"time"

	"gorm.io/gorm"
)

func main() {
	// validate environment variables
	env, err := config.Validate()
	if err != nil {
		log.Fatalf("error loading environment variables. err: %s", err.Error())
	}

	networkConfig := loadNetworkConfig()
	dbClient := getDBClient(env)
	provider := getProvider(env, networkConfig)
	internalNotification := getInternalNotification(env)
	// Database Stores
	addressStore := evm.NewAddressStore(dbClient)
	blockStore := evm.NewBlockStore(dbClient)
	transactionStore := evm.NewTransactionStore(dbClient)
	// Internal Service Communication
	// Listener control channels
	quitch := make(chan struct{})
	wg := &sync.WaitGroup{}
	defer close(quitch)

	// Creating network transaction app
	evmListener := evm.NewEvmListener(
		&evm.EvmListenerConfig{
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
	// Starting the app
	go evmListener.Start()

	// time.Sleep(time.Millisecond * 1)
	// evmListener.Stop()
	// fmt.Println("loopExitedGracefully")
	wg.Wait()
}

func loadNetworkConfig() evm.NetworkConfig {
	networkConfig, err := config.LoadNetworkConfig("network-config.json")
	if err != nil {
		log.Fatalf("[init] error loading network config file: %s", err.Error())

	}
	log.Println("[init] network config is loaded")
	return networkConfig

}

// // func loadEnvVariables() config.EnvVariables {
// // 	// Loading environment variables
// // 	env, err := config.LoadEnvVariableFile(".env")
// // 	if err != nil {
// // 		log.Fatalf("[init] error loading .env file: %s", err.Error())
// // 	}
// // 	log.Println("[init] environment variables are loaded")
// // 	return env
// // }

func getDBClient(env *config.Env) *gorm.DB {
	// Database connection
	dbClient, err := db.GetPostgresClient(db.DBClientSettings{
		DBUrl:             db.CreatePostgresDBUrl(env.POSTGRES_HOST, env.POSTGRES_PORT, env.POSTGRES_DBNAME, env.POSTGRES_USER, env.POSTGRES_PASSWORD),
		AutoMigrateModels: []any{&evm.Block{}, &evm.Transaction{}, &evm.Address{}},
	})
	if err != nil {
		log.Fatalf("[init] could not connect to the Postgres database: %s", err.Error())
	}
	log.Print("[init] connected to the database")
	return dbClient

}

func getProvider(env *config.Env, networkConfig evm.NetworkConfig) evm.EvmProvider {
	// Provider
	provider, err := evm.NewEvmProvider(env.PROVIDER_HOST, networkConfig.Network)
	if err != nil {
		log.Fatalf("[init] could not connect to the provider: %s", err.Error())
	}
	log.Print("[init] provider Initiated")
	return provider

}

func getInternalNotification(env *config.Env) eventbus.InternalNotification {
	internalNotification, err := eventbus.NewNatsInternalNotification(env.NATS_HOST)
	if err != nil {
		log.Fatalf("[init] could not connect to the nats service: %s", err.Error())
	}
	log.Print("[init] connected to the nats service")
	return internalNotification
}
