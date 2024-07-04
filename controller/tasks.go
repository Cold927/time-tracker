package controller

import (
	"github.com/gin-gonic/gin"
	"log"
)

// TaskCountdownStart Начать отсчет времени по задаче для пользователя
//
//	@Summary		Начать отсчет времени по задаче для пользователя
//	@Description	Начать отсчет времени по задаче для пользователя
//	@Tags			tasks
//	@Param			uid	query	string	true	"ID пользователя"
//	@Router			/tasks/countdown/start [post]
func TaskCountdownStart(c *gin.Context) {
	log.Println("Countdown Start")
}

// TaskCountdownEnd Закончить отсчет времени по задаче для пользователя
//
//	@Summary		Закончить отсчет времени по задаче для пользователя
//	@Description	Закончить отсчет времени по задаче для пользователя
//	@Tags			tasks
//	@Param			id	query	string	true	"ID задачи"
//	@Router			/tasks/countdown/end [patch]
func TaskCountdownEnd(c *gin.Context) {
	log.Println("Countdown End")
}
