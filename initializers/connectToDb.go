package initializers

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	// postgres://xcfuqlex:YFj5pRQUkFYoJb1-mHSv9MdWdb01ir7f@tiny.db.elephantsql.com/xcfuqlex

	var err error
	dsn := os.Getenv("DB")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

}
