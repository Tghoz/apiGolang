package controllers

import (
	"net/http"

	models "github.com/Tghoz/apiGolang/Model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	dataBase "github.com/Tghoz/apiGolang/DataBase"

	"golang.org/x/crypto/bcrypt"

	repo "github.com/Tghoz/apiGolang/Repository"
)

type login struct {
	Email    string
	Password string
}

func Register(c *gin.Context) {

	validate := validator.New()
	user := models.User{}

	db := dataBase.Db

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return

	}

	user.ID = uuid.New()
	user.Password = string(pass)

	result := repo.Create(db, &user)

	if result != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "failt to insert"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Create": true})
}

func Login(c *gin.Context) {

	validate := validator.New()
	body := login{}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := repo.FindOne(body.Email, body.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "credenciales invalidas"})
		return
	}

	if resp == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.Header("Access-Control-Expose-Headers", "x-access-token")
	c.Header("x-access-token", resp.Token)
	c.JSON(http.StatusOK, gin.H{"status": true})

}
