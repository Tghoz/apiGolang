package dataBase

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connection() {

	err := godotenv.Load(".env") 
	if err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}

	host := os.Getenv("host")
	port := os.Getenv("port")
	user := os.Getenv("user")
	password := os.Getenv("password")
	dbname := os.Getenv("dbname")

	var error error
	DSN := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, password, dbname)
	Db, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if error != nil {
		log.Fatal("( ͡ಠ ʖ̯ ͡ಠ) -> ", error)
	} else {
		log.Println("\n(👉ﾟヮﾟ)👍")
	}

}
