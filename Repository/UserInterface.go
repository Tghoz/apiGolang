package repository

import models "github.com/Tghoz/apiGolang/Model"

type UserInterface interface {
	Save(user models.User)
	Update(user models.User)
	Delete(userId int)
	FindById(userId int) (user models.User, err error)
	FindAll() []models.User
}
