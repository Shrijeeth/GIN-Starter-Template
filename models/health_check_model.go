package models

import (
	"gorm.io/gorm"
)

type HealthCheckModel struct {
	gorm.Model
	ID uint `gorm:"primaryKey,uniqueIndex"`
	Status string
}

func (HealthCheckModel) TableName() string {
	return "health_check"
}