package repository

import (
	"errors"

	dataBase "github.com/Tghoz/apiGolang/DataBase"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func FindAll[T any](model T, preloads ...string) ([]T, error) {
	var results []T
	tx := dataBase.Db
	for _, preload := range preloads {
		tx = tx.Preload(preload)
	}
	if err := tx.Order("created_at DESC").Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func FindById[T any](id string, model T, preloads ...string) (*T, error) {
	var result T
	tx := dataBase.Db
	for _, preload := range preloads {
		tx = tx.Preload(preload)
	}
	T_id, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	if err := tx.Limit(1).Where("id = ?", T_id).First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return &result, nil
}
