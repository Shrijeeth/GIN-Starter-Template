package config

import (
	"log"
	"mymodule/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(&models.HealthCheckModel{})
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	} else {
		log.Println("Database migrated successfully")
	}
}