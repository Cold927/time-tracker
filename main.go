package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	_ "time-tracker/docs"
	"time-tracker/handlers"
)

// @title Time Tracker API
// @version 1
// @Description Документация по сервису тайм трекера

// @host localhost:8500
// @BasePath /api/v1
func main() {
	router := gin.Default()
	people := router.Group("/users")
	tasks := router.Group("/tasks")

	user := people.Group("/users")
	{
		user.POST("/create", handlers.CreateUser)
		user.GET("/list")
		user.GET("/list/:id")

		user.DELETE("/delete")
	}

	task := tasks.Group("/tasks")
	{
		task.PATCH("/countdown/start")
		task.PATCH("/countdown/end")
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Fatal(router.Run(":8500")) // #TODO засунуть все переменные в .env
}
