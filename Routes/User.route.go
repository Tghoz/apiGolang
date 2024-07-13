package routes

import (
	controller "github.com/Tghoz/apiGolang/Controllers"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {

	router.GET("api/user", controller.GetUser)
	router.POST("api/user", controller.CreateUser)

}
