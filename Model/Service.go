package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Services struct {
	gorm.Model
	ID      uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name    string    `gorm:"not null" json:"name"`
	Price   float64   `gorm:"not null" json:"price"`
	Clients []Clients `gorm:"many2many:client_services;"`
}
