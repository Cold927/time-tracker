package controller

import (
	"github.com/gin-gonic/gin"
	"log"
)

// TaskCountdownStart Начать отсчет времени по задаче для пользователя
// @Summary Начать отсчет времени по задаче для пользователя
// @Description Начать отсчет времени по задаче для пользователя
// @Security bearerToken
// @Tags Tasks
// @Accept json
// @Produce json
// @Router /tasks/countdown/start [post]
func TaskCountdownStart(c *gin.Context) {
	log.Println("Countdown Start")
}

// TaskCountdownEnd Закончить отсчет времени по задаче для пользователя
// @Summary Закончить отсчет времени по задаче для пользователя
// @Description Закончить отсчет времени по задаче для пользователя
// @Security bearerToken
// @Tags Tasks
// @Accept json
// @Produce json
// @Router /tasks/countdown/end [patch]
func TaskCountdownEnd(c *gin.Context) {
	log.Println("Countdown End")
}
