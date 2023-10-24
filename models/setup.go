package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:Roxy@tcp(localhost:3306)/golang_db_token"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&customer{}, &track{})

	DB = database
}
