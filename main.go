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
	"time-tracker/docs"
	"time-tracker/model"
)

func main() {
	loadEnv()
	loadDatabase()
	setSwaggerInfo()
	serveApplication()
}

func setSwaggerInfo() {
	docs.SwaggerInfo.Title = "Time Tracker API"
	docs.SwaggerInfo.Description = "Документация по сервису тайм трекера"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:" + os.Getenv("APP_PORT")
	docs.SwaggerInfo.BasePath = "/api/v1"
}

/*
Подключение к базе данных и авто миграция структуры
*/
func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Task{})
}

/*
Подгрузка .env файла
*/
func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Ошибка при загрузке .env файла")
	}
}

/*
Запуск gin и назначение контроллеров
Порт запуска указывается в .env в поле APP_PORT
*/
func serveApplication() {
	router := gin.Default()
	v1 := router.Group("/api/v1")

	user := v1.Group("/users")
	{
		user.POST("/create", controller.CreateUser)
		user.PATCH("/update/:id", controller.UpdateUserData)
		user.GET("/info", controller.GetUsersInfo)
		user.GET("/find/:id", controller.GetUserById)
		user.DELETE("/delete/:id", controller.DeleteUser)
	}

	task := v1.Group("/tasks")
	{
		task.POST("/countdown/start/:uid", controller.TaskCountdownStart)
		task.PATCH("/countdown/end/:tid", controller.TaskCountdownEnd)
		task.GET("/info")
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Fatal(router.Run(":" + os.Getenv("APP_PORT")))
}
