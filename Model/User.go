package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	UserName string `gorm:"not null" json:"user_name" validate:"required" `
	Email    string `gorm:"not null;unique_index" json:"email" validate:"required,email"`
}
