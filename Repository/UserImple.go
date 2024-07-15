package repository

import (
	dto "github.com/Tghoz/apiGolang/Dto"
	models "github.com/Tghoz/apiGolang/Model"
	"gorm.io/gorm"

	"errors"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(Db *gorm.DB) UserInterface {
	return &UserRepositoryImpl{Db: Db}
}

func (u *UserRepositoryImpl) Delete(userId int) {
	var user models.User
	result := u.Db.Where("id = ?", userId).Unscoped().Delete(&user)

	if result == nil {
		panic(result)
	}
}

func (u *UserRepositoryImpl) FindAll() []models.User {
	var user []models.User
	result := u.Db.Find(&user)

	if result == nil {
		panic(result)
	}

	return user
}

func (u *UserRepositoryImpl) FindById(userId int) (user models.User, err error) {
	result := u.Db.Find(&user, userId)
	if result != nil {
		return user, nil
	} else {
		return user, errors.New("tag is not found")
	}
}

func (u *UserRepositoryImpl) Save(tags models.User) {
	result := u.Db.Create(&tags)

	if result == nil {
		panic(result)
	}
}

func (u *UserRepositoryImpl) Update(user models.User) {

	var updateTag = dto.UpdateTagsRequest{
		UserName: user.UserName,
		Email:    user.Email,
	}

	result := u.Db.Model(&user).Updates(updateTag)
	if result == nil {
		panic(result)
	}
}
