package DataBase

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DSN string = "host=localhost user=postgres dbname=test password=010300 port=5432"

var Db *gorm.DB

func Connection() {

	var error error

	Db, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if error != nil {
		log.Fatal("( ͡ಠ ʖ̯ ͡ಠ) -> ", error)
	} else {
		log.Println("\n(👉ﾟヮﾟ)👍")
	}

}
