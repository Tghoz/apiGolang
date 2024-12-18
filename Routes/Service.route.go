package routes

import (
	controller "github.com/Tghoz/apiGolang/Controllers/Client"

	"github.com/gin-gonic/gin"
)

func ServiceRouter(router *gin.Engine) {

	r := router.Group("/api/service")
	r.GET("", controller.GetService)
	r.POST("", controller.PostService)
	r.GET("/:id", controller.GetServiceByID)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
