package listeners

// import (
// 	"log"
// 	"open-payment-gateway/db"
// 	"open-payment-gateway/internal_notification"
// 	"open-payment-gateway/providers"
// 	"open-payment-gateway/types"
// 	"open-payment-gateway/utils"
// 	"sync"
// 	"testing"
// 	"time"
// )

// func TestEvmListenerStopFunction(t *testing.T) {
// 	// Loading environment variables
// 	env, err := utils.LoadEnvVariableFile("../.env")
// 	if err != nil {
// 		log.Fatalf("Error loading .env file: %s", err)
// 	}

// 	// Creating network object
// 	network := types.Network{
// 		Name:     env.NetworkName,
// 		Currency: env.NetworkCurrency,
// 		ChainID:  env.ChainID,
// 		Decimals: env.Decimals,
// 	}

// 	// Database connection
// 	dbClient, err := db.GetDBClient(db.DBClientSettings{
// 		DBUrl:             db.CreateDBUrl(env.DBUrl, env.DBPort, env.DBName, env.DBUser, env.DBPassword),
// 		AutoMigrateModels: []any{&types.Block{}, &types.Transaction{}, &types.Address{}},
// 	})
// 	if err != nil {
// 		panic("Could not connect to the Postgres database")
// 	}

// 	// Provider
// 	provider, err := providers.NewEvmProvider(env.ProviderUrl, network)
// 	if err != nil {
// 		panic("Could not connect to the provider")
// 	}

// 	// Database Stores
// 	addressStore := db.NewAddressStore(dbClient)
// 	blockStore := db.NewBlockStore(dbClient)
// 	transactionStore := db.NewTransactionStore(dbClient)

// 	// Internal Service Communication
// 	internalNotification := internal_notification.NewNatsInternalNotification(transactionStore)

// 	// Listener control channels
// 	quitch := make(chan struct{})
// 	wg := &sync.WaitGroup{}
// 	defer close(quitch)

// 	// Creating network transaction listener
// 	evmListener := NewEvmListener(
// 		&EvmListenerConfig{
// 			// Real Config
// 			Quitch: quitch,
// 			Wg:     wg,
// 			// Listener settings, also config
// 			Network:             network,
// 			StartingBlockNumber: env.StartingBlockNumber,
// 			// Stores
// 			AddressStore:     addressStore,
// 			BlockStore:       blockStore,
// 			TransactionStore: transactionStore,
// 			// Communication
// 			Notification: internalNotification,
// 			// Third Parties
// 			Provider:        provider,
// 			WaitForNewBlock: time.Second * 1,
// 		},
// 	)

// 	wg.Add(1)
// 	// Starting the listener
// 	go evmListener.Start()

// 	// time.Sleep(time.Millisecond * 1)
// 	didStop := evmListener.Stop()
// 	if !didStop {
// 		t.Fatal("The listener did not stop")
// 	}

// 	// fmt.Println("loopExitedGracefully")
// 	wg.Wait()
// }
