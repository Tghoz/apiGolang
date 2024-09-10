package dataBase

import (
	"log"
	"os"

	"fmt"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connection() {

	godotenv.Load(".env")
	// DSN := os.Getenv("DATABASE_PUBLIC_URL")

	username := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	databaseName := os.Getenv("DBNAME")
	databaseHost := os.Getenv("HOST")

	local := fmt.Sprintf("host=%s user=%s dbname=%s password=%s", databaseHost, username, databaseName, password)

	var error error

	Db, error = gorm.Open(postgres.Open(local), &gorm.Config{})

	if error != nil {
		log.Fatal("( ͡ಠ ʖ̯ ͡ಠ) -> ", error)
	} else {
		log.Println("\n(👉ﾟヮﾟ)👍")
	}

}
