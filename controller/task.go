package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time-tracker/model"
	"time-tracker/utils"
)

var task model.Task

// TaskCountdownStart Начать отсчет времени по задаче для пользователя
//
//	@Summary		Начать отсчет времени по задаче для пользователя
//	@Description	Начать отсчет времени по задаче для пользователя
//	@Tags			tasks
//	@Param			uid		path	uint		true	"ID пользователя"
//	@Param			task	body	model.Task	false	"Описание задачи"
//	@Router			/tasks/countdown/start/{uid} [post]
func TaskCountdownStart(c *gin.Context) {
	id := c.Param("uid")
	uid, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&task); err != nil {
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	task.UserID = uint(uid)
	newTask, err := task.CountdownStart()
	if err != nil {
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": newTask})
	log.Println("Countdown Start")
}

// TaskCountdownEnd Закончить отсчет времени по задаче для пользователя
//
//	@Summary		Закончить отсчет времени по задаче для пользователя
//	@Description	Закончить отсчет времени по задаче для пользователя
//	@Tags			tasks
//	@Param			tid	path	string	true	"ID задачи"
//	@Router			/tasks/countdown/end/{tid} [patch]
func TaskCountdownEnd(c *gin.Context) {
	tid := c.Param("tid")
	taskEnd, err := task.CountdownEnd(tid)
	if err != nil {
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, taskEnd)
	log.Println("Countdown End")
}
