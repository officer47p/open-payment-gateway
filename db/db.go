package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBClientSettings struct {
	DBUrl             string
	AutoMigrateModels []interface{}
}

func GetPostgresClient(s DBClientSettings) (*gorm.DB, error) {
	c, err := gorm.Open(postgres.Open(s.DBUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = c.AutoMigrate(s.AutoMigrateModels...)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func CreatePostgresDBUrl(url string, port int64, name string, user string, password string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, password, url, port, name)
}
