package Controller

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func connectForGorm() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/market?parseTime=true&loc=Asia%2FJakarta")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		return nil, err
	}
	db.LogMode(true)
	return db, err
}
