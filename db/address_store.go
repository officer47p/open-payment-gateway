package db

import (
	"log"
	"open-payment-gateway/types"

	"gorm.io/gorm"
)

type AddressStore interface {
	AddressExists(string) (bool, error)
}

type SQLAddressStore struct {
	client *gorm.DB
}

func NewAddressStore(c *gorm.DB) *SQLAddressStore {
	return &SQLAddressStore{client: c}
}

func (s *SQLAddressStore) AddressExists(a string) (bool, error) {
	var foundAddress string
	tx := s.client.Model(&types.Address{}).Select("address").Where("address ILIKE ?", a).Find(&foundAddress)
	if tx.Error != nil {
		log.Fatal(tx.Error)
	}
	if foundAddress == "" {
		return false, nil
	}
	return true, nil
}
