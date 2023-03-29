package db

import (
	"fmt"
	"vatansoft/pkg/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB //database

func init() {
	dsn := "root:12345678@tcp(127.0.0.1:3306)/vatansoft?charset=utf8mb4&parseTime=True&loc=Local"

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{PrepareStmt: true})
	if err != nil {
		//mongo
		fmt.Println(err.Error())
	}

	db = conn
	db.Debug().AutoMigrate(&model.Product{}) //Database migration
	db.Debug().AutoMigrate(&model.DeletedProduct{})
}

// returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
