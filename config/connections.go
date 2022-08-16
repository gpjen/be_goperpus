package config

import (
	"be_goperpus/books"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnDb() (db *gorm.DB) {
	dsn := "root:@tcp(127.0.0.1:3306)/goperpus?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&books.Books{})

	if err != nil {
		fmt.Println("error : database connection error")
	}
	fmt.Println("database connection success")
	return
}
