package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDataBase() {
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string
	fmt.Println(dbURI)

	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Print(err)
	}

	DB = conn
	DB.Debug().AutoMigrate(&User{}, &Task{}, &Tag{}) //Database migration
}
