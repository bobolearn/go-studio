package db

import (
	"fmt"
	model "go-studio/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (db *gorm.DB) {
	dsn := "root:@1781760347zzZZ@tcp(127.0.0.1:3306)/aq?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&model.User{}, &model.Question{}, &model.Answer{})
	fmt.Println("Database connection successful")
	return db
}
