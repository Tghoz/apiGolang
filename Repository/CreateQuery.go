package repository

import (

	dataBase "github.com/Tghoz/apiGolang/DataBase"
)

func Create[T any]( model T) error {
	db := dataBase.Db
	result := db.Create(&model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
