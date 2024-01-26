package config

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeDatabase() *gorm.DB {
	dsn := os.Getenv("DB_DSN")

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatalf("Error connecting database")
	}

	return db
}