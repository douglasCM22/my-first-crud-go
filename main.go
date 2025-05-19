package main

import (
	"log"
	"my-first-crud-go/src/configuration/logs"
	"my-first-crud-go/src/controller"
	"my-first-crud-go/src/controller/routes"
	"my-first-crud-go/src/service"
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

	userService := service.NewUserDomainService()
	userController := controller.NewUserController(userService)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	router.Run(":" + os.Getenv("PORT"))
}
