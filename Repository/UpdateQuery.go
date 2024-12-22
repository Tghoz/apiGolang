package repository

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// func [T any](db *gorm.DB, id string, model T) error {
// 	var result T
// 	tx := db
// 	T_id, err := uuid.Parse(id)
// 	if err != nil {
// 		return err
// 	}
// 	if err := tx.Limit(1).Where("id = ?", T_id).First(&result).Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return nil
// 		}
// 		return err
// 	}
// 	return tx.Model(&result).Updates(model).Error
// }

// Nueva función de actualización con preloads
func Update[T any](db *gorm.DB, id string, model T) error {
	var result T
	tx := db
	T_id, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	if err := tx.Limit(1).Where("id = ?", T_id).First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	return tx.Model(&result).Updates(model).Error
}
