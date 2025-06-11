package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() (*gorm.DB, error) {
	dsn := "host=localhost user=user password=pwd port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}
