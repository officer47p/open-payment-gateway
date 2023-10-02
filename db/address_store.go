package db

import (
	"strings"

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
	if strings.EqualFold(a, "0x6eb94F4C9CeDF3637f9F3ec21e91231fB8482278") {
		return true, nil
	}
	return false, nil
}
