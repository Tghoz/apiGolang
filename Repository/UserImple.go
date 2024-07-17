package repository

import (
	dataBase "github.com/Tghoz/apiGolang/DataBase"
	models "github.com/Tghoz/apiGolang/Model"
)

func Delete(id string) error {
	var user models.User
	dataBase.Db.Unscoped().Where("id = ?", id).Delete(&user)
	return nil
}

func FindById(id string) (*models.User, error) {

	var user models.User
	if err := dataBase.Db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func FindAll() ([]models.User, error) {
	var users []models.User
	if err := dataBase.Db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
