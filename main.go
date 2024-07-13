package main

import (
	"net/http"

	dataBase "github.com/Tghoz/apiGolang/DataBase"
	models "github.com/Tghoz/apiGolang/Model"
	routes "github.com/Tghoz/apiGolang/Routes"
	"github.com/gin-gonic/gin"
)

func main() {

	dataBase.Connection()

	dataBase.Db.AutoMigrate(models.User{})

	r := gin.Default()

	routes.UserRouter(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world from server Go.",
		})
	})

	r.Run(":3000")
}
