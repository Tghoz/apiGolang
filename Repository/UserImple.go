package repository

import (
	dataBase "github.com/Tghoz/apiGolang/DataBase"
	models "github.com/Tghoz/apiGolang/Model"
)

func Delete(id string) error {

	var user models.User
	result := dataBase.Db.Unscoped().Where("id = ?", id).Delete(&user)

	if result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}
