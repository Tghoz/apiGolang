package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// db.Connection()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world from server Go.",
		})
	})

	r.Run(":3000")
}
