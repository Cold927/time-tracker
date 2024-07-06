package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time-tracker/model"
	"time-tracker/utils"
)

var task model.Task

// TaskCountdownStart Начать отсчет времени по задаче для пользователя
//
//	@Summary		Начать отсчет времени по задаче для пользователя
//	@Description	Начать отсчет времени по задаче для пользователя
//	@Tags			Задачи
//	@Accept			json
//	@Produce		json
//	@Param			uid		path		string				true	"ID пользователя"
//	@Param			task	body		model.TaskCreate	false	"Описание задачи"
//	@Success		200		{object}	utils.HTTPSuccess
//	@Failure		400		{object}	utils.HTTPError
//	@Failure		404		{object}	utils.HTTPError
//	@Router			/api/v1/tasks/countdown/start/{uid} [post]
func TaskCountdownStart(c *gin.Context) {
	id := c.Param("uid")
	uid, err := uuid.Parse(id)
	if err != nil {
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&task); err != nil {
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	task.UserID = uid
	newTask, err := task.CountdownStart()
	if err != nil {
		utils.NewError(c, http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": newTask})
}

// TaskCountdownEnd Закончить отсчет времени по задаче для пользователя
//
//	@Summary		Закончить отсчет времени по задаче для пользователя
//	@Description	Закончить отсчет времени по задаче для пользователя
//	@Tags			Задачи
//	@Accept			json
//	@Produce		json
//	@Param			tid	path		string	true	"ID задачи"
//	@Success		200	{object}	utils.HTTPSuccess
//	@Failure		400	{object}	utils.HTTPError
//	@Failure		404	{object}	utils.HTTPError
//	@Router			/api/v1/tasks/countdown/end/{tid} [patch]
func TaskCountdownEnd(c *gin.Context) {
	tid := c.Param("tid")
	_, err := task.CountdownEnd(tid)
	if err != nil {
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Задача обновлена"})
}

// TasksInfo Получение трудозатрат по пользователю
//
//	@Summary		Получение трудозатрат по пользователю
//	@Description	Получение трудозатрат по пользователю за период задача-сумма часов и минут
//	@Tags			Задачи
//	@Accept			json
//	@Produce		json
//	@Param			uid			path		string	true	"ID пользователя"
//	@Param			startDate	query		string	true	"Начальная дата"	default(2024-07-01T00:00:00)
//	@Param			endDate		query		string	true	"Конечная дата"		default(2024-07-01T23:59:59)
//	@Success		200			{array}		model.TaskResponse
//	@Failure		400			{object}	utils.HTTPError
//	@Failure		404			{object}	utils.HTTPError
//	@Router			/api/v1/tasks/info/{uid} [get]
func TasksInfo(c *gin.Context) {
	uid := c.Param("uid")
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")
	tasksInfo, err := task.Info(uid, startDate, endDate)
	if err != nil {
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, tasksInfo)
	log.Println("Информация о трудозатратах")
}
