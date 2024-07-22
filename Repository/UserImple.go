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
	if err := dataBase.Db.Limit(1).Find(&user, id).Error; err != nil {
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

func Create(user *models.User) (err error) {
	result := dataBase.Db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return
}

func Update(user *models.User, body models.User) error {

	result := dataBase.Db.Model(&user).Updates(body)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
