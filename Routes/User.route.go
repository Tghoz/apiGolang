package routes

import (
	controller "github.com/Tghoz/apiGolang/Controllers"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {

	routes := router.Group("api/user")

	routes.GET("", controller.GetUser)
	routes.POST("", controller.CreateUser)
	routes.GET("/:id", controller.GetUserByID)
	routes.DELETE("/:id", controller.DeleteUser)
	//routes.PUT("/:id", controller.UpdateUser)
}
