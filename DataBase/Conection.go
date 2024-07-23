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

	godotenv.Load(".env")

	DSN := os.Getenv("DATABASE_URL")
	fmt.Print(DSN)

	var error error

	Db, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if error != nil {
		log.Fatal("( ͡ಠ ʖ̯ ͡ಠ) -> ", error)
	} else {
		log.Println("\n(👉ﾟヮﾟ)👍")
	}

}
