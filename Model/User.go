package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserName string    `gorm:"not null" json:"user_name" validate:"required" `
	Password string    `gorm:"not null" json:"password" validate:"required" `
	Email    string    `gorm:"not null;unique_index" json:"email" validate:"required,email"  crypto:"aes"`
}




