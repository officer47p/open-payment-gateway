package main

import (
	"log"
	"open-payment-gateway/pkg/config"
	"open-payment-gateway/pkg/db"
	"open-payment-gateway/pkg/evm"

	"golang.org/x/exp/rand"
	"gorm.io/gorm"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func main() {
	// validate environment variables
	env, err := config.Validate()
	if err != nil {
		log.Fatalf("error loading environment variables. err: %s", err.Error())
	}

	seedAddresses := []string{
		"0x07C2e3A3c6a3d2efAce6A1829A8257913B36e942", // 9797391
		"0xB2EeA5319Ea10fd1A426919483ffC7CbFD7430a2", // 9797391
		"0x01d01c0988213E493c690E5088eF8A8ef23Fe6f5", // 9797393
	}

	dbClient := getDBClient(env)
	// Database Stores
	addressStore := evm.NewAddressStore(dbClient)

	for _, address := range seedAddresses {
		err = addressStore.InsertAddress(&evm.Address{
			Address: address,
			HDPath:  randSeq(10),
		})

		if err != nil {
			log.Fatalf("error inserting address: %s", err.Error())
		}
	}

}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

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
