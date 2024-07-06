package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"time-tracker/config"
	"time-tracker/controller"
	"time-tracker/database"
	"time-tracker/database/migration"
	"time-tracker/docs"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal("Невозможно получить конфигурацию: ", err)
	}
	loadDatabase(cfg)
	setSwaggerInfo(cfg)
	serveApplication(cfg)
}

func setSwaggerInfo(cfg config.Config) {
	docs.SwaggerInfo.Title = "Time Tracker API"
	docs.SwaggerInfo.Description = "Документация по сервису тайм трекера"
	docs.SwaggerInfo.Version = "0.1.0"
	docs.SwaggerInfo.Host = "localhost:" + cfg.Port
	docs.SwaggerInfo.BasePath = "/"
}

/*
Подключение к базе данных и авто миграция структуры
*/
func loadDatabase(cfg config.Config) {
	database.Connect(cfg)
	migration.StartUserMigration(database.Database)
	migration.StartTaskMigration(database.Database)
}

/*
Запуск gin и назначение контроллеров
Порт запуска указывается в .env в поле APP_PORT
*/
func serveApplication(cfg config.Config) {
	gin.SetMode(cfg.RunMode)
	router := gin.Default()
	router.Use(Cors(cfg))
	v1 := router.Group("/api/v1")

	user := v1.Group("/users")
	{
		user.POST("/create", controller.CreateUser)
		user.PATCH("/update/:id", controller.UpdateUserData)
		user.GET("/list", controller.GetUsersList)
		user.GET("/info", controller.GetUserInfo)
		user.GET("/find/:id", controller.GetUserById)
		user.DELETE("/delete/:id", controller.DeleteUser)
	}

	task := v1.Group("/tasks")
	{
		task.POST("/countdown/start/:uid", controller.TaskCountdownStart)
		task.PATCH("/countdown/end/:tid", controller.TaskCountdownEnd)
		task.GET("/info/:uid", controller.TasksInfo)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Fatal(router.Run(":" + cfg.Port))
}

func Cors(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", cfg.Cors.AllowOrigins)
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE. UPDATE")
		c.Header("content-type", "application/json")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
