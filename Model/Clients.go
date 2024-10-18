package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Clients struct {
	gorm.Model
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey"`
	Name      string     `gorm:"not null" json:"name"`
	Telephone string     `gorm:"not null" json:"telephone"`
	Status    string     `gorm:"not null" json:"status"`
	Services  []Services `gorm:"many2many:client_services;"`
	History   []Payments `gorm:"foreignKey:ClientID"`
}
