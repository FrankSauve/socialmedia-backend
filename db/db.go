package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres() *gorm.DB {
	dsn := "host=localhost user=postgres password=Fcs8803! dbname=socialmedia port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	log.Println("Connected to Postgres!")

	return db
}
