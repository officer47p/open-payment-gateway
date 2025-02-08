package evm

import (
	"testing"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func TestDatabaseOperations(t *testing.T) {
	// Open a SQLite in-memory database for testing purposes
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error opening the database: %v", err)
	}
	// defer db.Close()

	// Migrate the tables
	err = db.AutoMigrate(&Block{}, &Transaction{}, &Address{})
	if err != nil {
		t.Fatalf("Error migrating tables: %v", err)
	}

	// Test inserting and querying a Block record
	block := Block{
		Network:           "TestNetwork",
		BlockNumber:       1,
		BlockHash:         "BlockHash1",
		PreviousBlockHash: "PreviousBlockHash1",
	}
	db.Create(&block)

	var retrievedBlock Block
	db.First(&retrievedBlock, block.ID)
	if retrievedBlock.Network != block.Network || retrievedBlock.BlockNumber != block.BlockNumber {
		t.Errorf("Retrieved Block does not match the inserted Block")
	}

	// Test inserting and querying a Transaction record
	transaction := Transaction{
		Broadcasted: true,
		BlockNumber: 1,
		BlockHash:   "BlockHash1",
		Network:     "TestNetwork",
		Currency:    "ETH",
		TxHash:      "TxHash1",
		TxType:      "Transfer",
		Value:       "10.5",
		From:        "FromAddress",
		To:          "ToAddress",
	}
	db.Create(&transaction)

	var retrievedTransaction Transaction
	db.First(&retrievedTransaction, transaction.ID)
	if retrievedTransaction.TxHash != transaction.TxHash || retrievedTransaction.TxType != transaction.TxType {
		t.Errorf("Retrieved Transaction does not match the inserted Transaction")
	}

	// Test inserting and querying an Address record
	address := Address{
		Address: "TestAddress",
		HDPath:  "m/44'/60'/0'/0/0",
	}
	db.Create(&address)

	var retrievedAddress Address
	db.First(&retrievedAddress, address.ID)
	if retrievedAddress.Address != address.Address || retrievedAddress.HDPath != address.HDPath {
		t.Errorf("Retrieved Address does not match the inserted Address")
	}
}
