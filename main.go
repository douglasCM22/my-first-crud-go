package main

import (
	"log"
	"my-first-crud-go/src/configuration/logs"
	"my-first-crud-go/src/controller/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	logs.Info("Starting application...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	router.Run(":" + os.Getenv("PORT"))
}
