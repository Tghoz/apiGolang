package controller

import (
	dataBase "github.com/Tghoz/apiGolang/DataBase"
	models "github.com/Tghoz/apiGolang/Model"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {

	var user []models.User
	dataBase.Db.Find(&user)
	c.JSON(200, &user)

}

func CreateUser(c *gin.Context) {

	body := models.User{}
	c.BindJSON(&body)
	user := &models.User{UserName: body.UserName, Email: body.Email}
	result := dataBase.Db.Create(&user)
	if result.Error != nil {
		c.JSON(500, gin.H{"Error": "failt to insert"})
		return
	}
	c.JSON(200, result)
}
