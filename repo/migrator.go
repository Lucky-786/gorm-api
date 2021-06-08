package repo

import (
	"Assignemnts/APIs/models"
	"fmt"
	"log"

	//"github.com/lucky-786/gorm-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitialMigration() (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlite.Open("./data/database.db"), &gorm.Config{})
	DB = db
	if err != nil {
		fmt.Println(err.Error())
		panic("Faild to connect to database")
	} else {
		log.Println("Connection established")
	}
	x := db.Migrator().HasTable(&models.User{})
	if !x {
		db.AutoMigrate(&models.User{})
	}
	return db, nil
}
