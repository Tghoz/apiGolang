package main

import (
	dataBase "github.com/Tghoz/apiGolang/DataBase"
	models "github.com/Tghoz/apiGolang/Model"
	routes "github.com/Tghoz/apiGolang/Routes"
	"github.com/gin-gonic/gin"
)

func main() {

	dataBase.Connection()

	modelsMigrate := []interface{}{
		&models.User{},
		&models.Clients{},
		&models.Services{},
		&models.Payments{},
	}

	dataBase.Db.AutoMigrate(modelsMigrate...)

	r := gin.Default()
	routes.UserRouter(r)
	routes.ClientRouter(r)
	routes.ServiceRouter(r)

	
	r.Run(":3000")
}
