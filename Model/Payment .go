package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payments struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	ClientID uuid.UUID `gorm:"index" json:"client_id"`
	Amount   float64   `gorm:"not null" json:"amount"`
	Type     string    `gorm:"not null" json:"type"`
	Date     time.Time `gorm:"not null" json:"date"`
}
