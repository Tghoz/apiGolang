package dto

import (
	models "github.com/Tghoz/apiGolang/Model"
)

type UserDto struct {
	ID       string
	UserName string
	Email    string
	Password string
}

func UserDtoMap(u models.User) UserDto {
	return UserDto{
		ID:       u.ID.String(),
		UserName: u.UserName,
		Email:    u.Email,
		Password: u.Password,
	}
}
