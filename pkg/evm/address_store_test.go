package evm

// import (
// 	"log"
// 	"open-payment-gateway/types"
// 	"testing"

// 	"github.com/glebarez/sqlite"
// 	"gorm.io/gorm"
// )

// func setupTestDB() (*gorm.DB, func()) {
// 	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Failed to open database: %v", err)
// 	}

// 	err = db.AutoMigrate(&types.Address{})
// 	if err != nil {
// 		log.Fatalf("Failed to auto-migrate: %v", err)
// 	}

// 	return db, func() {
// 		db.Migrator().DropTable(&types.Address{})
// 	}
// }

// func TestAddressStore_AddressExists(t *testing.T) {
// 	db, cleanup := setupTestDB()
// 	defer cleanup()

// 	addressStore := NewAddressStore(db)

// 	// Test for an address that does not exist in the database
// 	exists, err := addressStore.AddressExists("0xAddressNotInDB")
// 	if err != nil {
// 		t.Errorf("AddressExists returned an error: %v", err)
// 	}
// 	if exists {
// 		t.Errorf("Expected address not to exist in the database, but it does")
// 	}

// 	// Test for an address that exists in the database
// 	newAddress := types.Address{
// 		Address: "0xAddressInDB",
// 		HDPath:  "m/44'/60'/0'/0/0",
// 	}
// 	db.Create(&newAddress)

// 	exists, err = addressStore.AddressExists("0xAddressInDB")
// 	if err != nil {
// 		t.Errorf("AddressExists returned an error: %v", err)
// 	}
// 	if !exists {
// 		t.Errorf("Expected address to exist in the database, but it does not")
// 	}
// }
