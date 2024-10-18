package routes

import (
	"time"

	controller "github.com/Tghoz/apiGolang/Controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	middleware "github.com/Tghoz/apiGolang/Middleware"
)

func UserRouter(router *gin.Engine) {

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{`http://localhost:4321`}

	config.AllowMethods = []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	router.Use(cors.New(config))

	authGroup := router.Group("/api/user")
	authGroup.Use(middleware.JwtVerify())
	{
		authGroup.GET("", controller.GetUser)
		authGroup.GET("/:id", controller.GetUserByID)
		authGroup.DELETE("/:id", controller.DeleteUser)
		authGroup.PUT("/:id", controller.UpdateUser)
	}

	// Grupo para las rutas de autenticación
	aut := router.Group("/api/auth")
	{
		aut.POST("/login", controller.Login)
		aut.POST("/register", controller.Register)
	}

	// Grupo para las rutas de autenticación con Google
	autGoogle := router.Group("/api/auth/google")
	{
		autGoogle.GET("", controller.GoogleLogin)
		autGoogle.GET("/redirect", controller.GoogleRedirect)

	}

}
