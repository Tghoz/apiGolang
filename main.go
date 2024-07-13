package main

import (
	"net/http"

	DataBase "github.com/Tghoz/apiGolang/DataBase"
	models "github.com/Tghoz/apiGolang/Model"
	"github.com/gin-gonic/gin"
)

func main() {

	DataBase.Connection()

	DataBase.Db.AutoMigrate(models.User{})

	r := gin.Default()


	
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world from server Go.",
		})
	})

	r.Run(":3000")
}
