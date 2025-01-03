package routes

import (
	controller "github.com/Tghoz/apiGolang/Controllers/Client"

	"github.com/gin-gonic/gin"
)

func PaymentRouter(router *gin.Engine) {

	r := router.Group("/api/payment")
	r.GET("", controller.GetPayment)
	r.POST("", controller.PostPayment)
	r.GET("/:id", controller.GetPaymentByID)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong payment",
		})
	})
}
