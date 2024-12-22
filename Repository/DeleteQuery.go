package repository

import (
	"errors"

	dataBase "github.com/Tghoz/apiGolang/DataBase"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Delete[T any](id string, model T) error {
	var result T
	tx := dataBase.Db
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
	return tx.Unscoped().Delete(&result).Error
}
