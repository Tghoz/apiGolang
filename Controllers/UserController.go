package controller

import (
	"net/http"

	models "github.com/Tghoz/apiGolang/Model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	repo "github.com/Tghoz/apiGolang/Repository"
)

func GetUser(c *gin.Context) {

	user, err := repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
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
	result := repo.Create(user)

	if result != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "failt to insert"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Create": true})
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, _ := repo.FindById(id)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	if user != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {

	id := c.Param("id")
	_, err := repo.FindById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	repo.Delete(id)
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

	data := &models.User{UserName: body.UserName, Email: body.Email}
	result := repo.Update(user, *data)

	if result != nil {
		c.JSON(500, gin.H{"error": true, "message": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, &user)

}
