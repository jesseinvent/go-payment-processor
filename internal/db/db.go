package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func ConnectDB (dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{});

	if err != nil {
		return nil, err
	}

	log.Println("Database successfully connected..")

	return db, nil
}