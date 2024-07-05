package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time-tracker/model"
	"time-tracker/utils"
)

type userResponse struct {
	Data []model.User `json:"data"`
}

var user model.User
var userCreate model.UserCreate

// CreateUser Создает нового пользователя
//
//	@Summary		Создает нового пользователя
//	@Description	Создает нового пользователя
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		model.UserCreate	true	"Новый пользователь"
//	@Success		201		{object}	userResponse
//	@Failure		400		{object}	utils.HTTPError
//	@Router			/users/create [post]
func CreateUser(c *gin.Context) {
	if err := c.ShouldBindJSON(&userCreate); err != nil {
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	savedUser, err := user.Save(&userCreate)
	if err != nil {
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": savedUser})
	log.Println("Пользователь был удачно создан")
}

// UpdateUserData Изменение данных пользователя
//
//	@Summary		Изменение данных пользователя
//	@Description	Изменение данных пользователя
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string				true	"Идентификатор пользователя"
//	@Param			user	body		model.UserCreate	true	"Изменение данных пользователя"
//	@Success		200		{object}	userResponse
//	@Failure		400		{object}	utils.HTTPError
//	@Router			/users/update/{id} [patch]
func UpdateUserData(c *gin.Context) {
	id := c.Param("id")
	if err := c.ShouldBindJSON(&userCreate); err != nil {
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}
	updatedUser, err := user.UpdateData(id, &userCreate)
	if err != nil {
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}
	fmt.Println("Обновление данных о пользователе")
	c.JSON(http.StatusOK, updatedUser)
}

// GetUsersInfo Получение данных о всех пользователях
//
//	@Summary		Получение данных о всех пользователях
//	@Description	Получение данных о всех пользователях
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			page			query		string	false	"Укажите с какой страницы смотреть"	default(1)
//	@Param			limit			query		string	false	"Укажите какое количество выводить"	default(10)
//	@Param			sort			query		string	false	"Сортировать данные"				example(asc, desc)
//	@Param			field			query		string	false	"Поле для сортировки"				example(Id, Surname, Name, Patronymic, Address, passport_series, passport_number)
//	@Param			passportSeries	query		int		false	"Поиск по серии паспорта"
//	@Param			passportNumber	query		int		false	"Поиск по номеру паспорта"
//	@Param			search			query		string	false	"Поиск по полям"
//	@Success		200				{array}		userResponse
//	@Failure		400				{object}	utils.HTTPError
//	@Router			/users/info [get]
func GetUsersInfo(c *gin.Context) {
	var pagination utils.Pagination
	sort := c.DefaultQuery("field", "Id") + " " + c.DefaultQuery("sort", "desc")
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}
	pagination.Page = page
	pagination.Limit = limit
	pagination.Sort = sort

	listInfo, err := user.ListUsers(pagination)
	if err != nil {
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, listInfo)
	fmt.Println("Получение данных о всех пользователях")
}

// GetUserById Получение данных о пользователе
//
//	@Summary		Получение данных о пользователе по ID
//	@Description	Получение данных о пользователе по ID
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Идентификатор пользователя"
//	@Success		200	{object}	userResponse
//	@Failure		404	{object}	utils.HTTPError
//	@Router			/users/find/{id} [get]
func GetUserById(c *gin.Context) {
	id := c.Param("id")

	foundUser, err := user.FindUserById(id)
	if err != nil {
		utils.NewError(c, http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, foundUser)
}

// DeleteUser Удаление пользователя
//
//	@Summary		Удаление пользователя
//	@Description	Изменение данных пользователя
//	@Tags			users
//	@Produce		json
//	@Success		200	{object}	utils.HTTPSuccess
//	@Failure		404	{object}	utils.HTTPError
//	@Param			id	path		string	true	"Идентификатор пользователя"
//	@Router			/users/delete/{id} [delete]
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	deletedUser, err := user.DeleteUser(id)
	if err != nil {
		utils.NewError(c, http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, deletedUser)
	fmt.Println("Пользователь удален")
}
