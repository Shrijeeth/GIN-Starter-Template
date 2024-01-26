package config

import (
	"gorm.io/gorm"
)

type Config struct {
	Db *gorm.DB
}

func InitializeConfigs() Config {
	Db := InitializeDatabase()
	RunMigrations(Db)

	config := Config{
		Db: Db,
	}
	return config
}