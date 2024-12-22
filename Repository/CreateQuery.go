package repository

import (
	"gorm.io/gorm"
)

func Create[T any](db *gorm.DB, model T) error {
	result := db.Create(&model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
