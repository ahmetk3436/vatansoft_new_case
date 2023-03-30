package db

import (
	"fmt"
	"vatansoft/pkg/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB //database

func init() {
	dsn := "root:toor@tcp(45.12.81.218:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{PrepareStmt: true})
	if err != nil {
		//mongo
		fmt.Println(err.Error())
	}

	db = conn
	db.AutoMigrate(&model.Product{}) //Database migration
	db.AutoMigrate(&model.Categories{})
}

// returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
