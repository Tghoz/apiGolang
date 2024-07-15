package dataBase

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DSN string = "host=localhost user=postgres dbname=test password=010300 port=5432"

func Connection() *gorm.DB {

	db, error := gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if error != nil {
		log.Fatal("( ͡ಠ ʖ̯ ͡ಠ) -> ", error)
	} else {
		log.Println("\n(👉ﾟヮﾟ)👍")
	}
	return db

}
