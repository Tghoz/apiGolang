package routes

import (
	controller "github.com/Tghoz/apiGolang/Controllers"

	"github.com/gin-gonic/gin"
)

func ClientRouter(router *gin.Engine) {

	r := router.Group("/api/client")
	r.GET("", controller.GetClient)
	r.POST("", controller.PostClient)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
