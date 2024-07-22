package dto

import models "github.com/Tghoz/apiGolang/Model"

type UserDto struct {
	ID       uint
	UserName string
	Email    string
}

func UserDtoMap(u models.User) UserDto {

	return UserDto{
		ID:       u.ID,
		UserName: u.UserName,
		Email:    u.Email,
	}
}
