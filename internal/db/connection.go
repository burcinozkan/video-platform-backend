package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// DSN = Data Source Name
	dsn := "host=localhost user=admin password=password dbname=video_platform port=5432 sslmode=disable"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(" Failed to connect to database: ", err)
	}

	DB = database
	fmt.Println(" Connected to PostgreSQL successfully!")
}
