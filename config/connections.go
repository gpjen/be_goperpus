package config

import (
	"be_goperpus/books"
	"be_goperpus/users"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnDb is to connect your databases
func ConnDb() (db *gorm.DB) {

	err := godotenv.Load()
	if err != nil {
		panic("failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbDatabse := os.Getenv("DB_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbDatabse)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("database connection error")
	}

	db.AutoMigrate(&books.Books{})
	db.AutoMigrate(&users.Users{})

	fmt.Println("database connection success")
	return
}

// CloseDb is function to close your connection betwen app and db
func CloseDb(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("failed to close conection from database")
	}
	dbSQL.Close()
}
