package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	UserName string `gorm:"not null"`
	Email    string `gorm:"not null;unique_index"`

}