package app

import (
	"log"
	"os"
	"task-api/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}
}

var router = gin.Default()

func StartApp() {
	models.ConnectDataBase()

	route()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080" //localhost
	}
	log.Fatal(router.Run(":" + port))
}
