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

	// postgresDb, err := db.DB()

	// if err != nil {
	// 	return nil, err
	// }

	// // Set connection pool settings
	// postgresDb.SetMaxOpenConns(10)

	// postgresDb.SetMaxIdleConns(5)

	// postgresDb.SetConnMaxLifetime(5 * time.Minute)
	

	log.Println("Database successfully connected..")

	return db, nil
}