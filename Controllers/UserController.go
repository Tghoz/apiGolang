package controllers

import (
	"net/http"

	dto "github.com/Tghoz/apiGolang/Dto"
	models "github.com/Tghoz/apiGolang/Model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"golang.org/x/crypto/bcrypt"

	repo "github.com/Tghoz/apiGolang/Repository"
)


func GetUser(c *gin.Context) {
	user, err := repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var userDto []dto.UserDto
	for _, u := range user {
		userDto = append(userDto, dto.UserDtoMap(u))
	}

	c.JSON(http.StatusOK, &userDto)
}


func GetUserByID(c *gin.Context) {

	id := c.Param("id")
	user, err := repo.FindById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Usuario no encontrado"})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Usuario no encontrado"})
		return
	}

	userDto := dto.UserDtoMap(*user)

	c.JSON(http.StatusOK, &userDto)
}

func DeleteUser(c *gin.Context) {

	id := c.Param("id")
	user, err := repo.FindById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user.UserName == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := repo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"delete": true})

}

func UpdateUser(c *gin.Context) {

	id := c.Param("id")
	user, err := repo.FindById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	body := models.User{}
	validate := validator.New()
	c.BindJSON(&body)

	if err := validate.Struct(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	body.Password = string(pass)

	result := repo.Update(user, body)

	if result != nil {
		c.JSON(500, gin.H{"error": true, "message": "Failed to update user"})
		return
	}

	userDto := dto.UserDtoMap(*user)
	c.JSON(http.StatusOK, &userDto)

}
