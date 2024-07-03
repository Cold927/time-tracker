package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
	"time-tracker/controller"
	"time-tracker/database"
	_ "time-tracker/docs"
	"time-tracker/model"
)

// @title Time Tracker API
// @version 1
// @Description Документация по сервису тайм трекера

// @host localhost:8500
// @BasePath /api/v1
func main() {
	loadEnv()
	//loadDatabase()
	serveApplication()
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Task{})
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func serveApplication() {
	router := gin.Default()
	people := router.Group("/users")
	tasks := router.Group("/tasks")

	user := people.Group("/users")
	{
		user.POST("/create", controller.CreateUser)
		user.PATCH("/change", controller.ChangeUserData)
		user.GET("/list", controller.GetUsersList)
		user.GET("/list/:id", controller.GetUserById)

		user.DELETE("/delete", controller.DeleteUser)
	}

	task := tasks.Group("/tasks")
	{
		task.POST("/countdown/start", controller.TaskCountdownStart)
		task.PATCH("/countdown/end", controller.TaskCountdownEnd)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Fatal(router.Run(":" + os.Getenv("APP_PORT")))
}
