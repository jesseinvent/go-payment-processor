package db

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func ConnectDB (dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{});

	if err != nil {
		return nil, err
	}

	postgresDb, err := db.DB()

	if err != nil {
		return nil, err
	}

	// Set connection pool settings
	postgresDb.SetMaxOpenConns(25) // max open connections to the DB
	postgresDb.SetMaxIdleConns(10) // max idle connections kept in pool
	postgresDb.SetConnMaxLifetime(5 * time.Minute) // max time a connection can be reused
	
	stats := postgresDb.Stats()

	log.Printf("DB Stats: \n Open: %d, Idle: %d, InUse: %d", 
		stats.OpenConnections,
		stats.Idle,
		stats.InUse,
	)

	log.Println("Database successfully connected..")

	return db, nil
}