package controller

import (
	"net/http"

	dataBase "github.com/Tghoz/apiGolang/DataBase"
	models "github.com/Tghoz/apiGolang/Model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetUser(c *gin.Context) {
	var user []models.User
	// dataBase.Db.Find(&user)
	c.JSON(http.StatusOK, &user)
}

func CreateUser(c *gin.Context) {

	validate := validator.New()
	body := models.User{}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &models.User{UserName: body.UserName, Email: body.Email}
	result := dataBase.Db.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "failt to insert"})
		return
	}

	respont := models.User{UserName: body.UserName, Email: body.Email}
	c.JSON(http.StatusOK, respont)
}

func GetUserByID(c *gin.Context) {

	id := c.Param("id")
	var user models.User
	dataBase.Db.First(&user, id)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "user not found"})
		return
	}

	c.JSON(200, &user)

}

func DeleteUser(c *gin.Context) {

	id := c.Param("id")

	var user models.User
	result := dataBase.Db.Unscoped().Delete(&user, id)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"delete": true})

}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	dataBase.Db.First(&user, id)

	body := models.User{}

	c.BindJSON(&body)
	data := &models.User{UserName: body.UserName, Email: body.Email}
	result := dataBase.Db.Model(&user).Updates(data)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": true, "message": "Failed to update"})
		return
	}

	c.JSON(200, &user)

}
