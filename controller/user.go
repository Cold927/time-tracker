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
//	@Tags			Пользователи
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
	log.Println("Пользователь был успешно создан")
}

// UpdateUserData Изменение данных пользователя
//
//	@Summary		Изменение данных пользователя
//	@Description	Изменение данных пользователя
//	@Tags			Пользователи
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string				true	"Идентификатор пользователя"
//	@Param			user	body		model.UserCreate	true	"Изменение данных пользователя"
//	@Success		200		{object}	utils.HTTPSuccess
//	@Failure		400		{object}	utils.HTTPError
//	@Router			/api/v1/users/update/{id} [patch]
func UpdateUserData(c *gin.Context) {
	id := c.Param("id")
	if err := c.ShouldBindJSON(&userCreate); err != nil {
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}
	_, err := user.UpdateData(id, &userCreate)
	if err != nil {
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}
	log.Println("Обновление данных о пользователе")
	c.JSON(http.StatusOK, gin.H{"message": "Пользователь успешно обновлен"})
}

// GetUsersList Получение данных о всех пользователях
//
//	@Summary		Получение данных о всех пользователях
//	@Description	Получение данных о всех пользователях
//	@Tags			Пользователи
//	@Accept			json
//	@Produce		json
//	@Param			page	query		string	false	"Укажите с какой страницы смотреть"	default(1)
//	@Param			limit	query		string	false	"Укажите какое количество выводить"	default(10)
//	@Param			sort	query		string	false	"Сортировать данные"				example(asc, desc)
//	@Param			field	query		string	false	"Поле для сортировки"				example(Id, Surname, Name, Patronymic, Address, PassportSeries, PassportNumber)
//	@Param			search	query		string	false	"Поиск по полям"
//	@Success		200		{object}	utils.Pagination
//	@Failure		400		{object}	utils.HTTPError
//	@Router			/api/v1/users/list [get]
func GetUsersList(c *gin.Context) {
	var pagination utils.Pagination
	field := utils.ToFormatCase(c.DefaultQuery("field", "Id"))
	search := c.DefaultQuery("search", "")
	sort := field + " " + c.DefaultQuery("sort", "desc")
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

	listInfo, err := user.ListUsers(pagination, field, search)
	if err != nil {
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, listInfo)
	log.Println("Получение данных о всех пользователях")
}

// GetUserInfo Получение данных о пользователе по паспорту
//
//	@Summary		Получение данных о пользователе по паспорту
//	@Description	Получение данных о пользователе по серии и номеру паспорта
//	@Tags			Пользователи
//	@Accept			json
//	@Produce		json
//	@Param			passportSeries	query		int	true	"Поиск по серии паспорта"
//	@Param			passportNumber	query		int	true	"Поиск по номеру паспорта"
//	@Success		200				{object}	userResponse
//	@Failure		400				{object}	utils.HTTPError
//	@Failure		404				{object}	utils.HTTPError
//	@Router			/api/v1/users/info [get]
func GetUserInfo(c *gin.Context) {
	series, _ := strconv.Atoi(c.Query("passportSeries"))
	number, _ := strconv.Atoi(c.Query("passportNumber"))

	userInfo, err := user.Info(series, number)
	if err != nil {
		utils.NewError(c, http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, userInfo)
}

// GetUserById Получение данных о пользователе
//
//	@Summary		Получение данных о пользователе по ID
//	@Description	Получение данных о пользователе по ID
//	@Tags			Пользователи
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"ID пользователя"
//	@Success		200	{object}	userResponse
//	@Failure		404	{object}	utils.HTTPError
//	@Router			/api/v1/users/find/{id} [get]
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
//	@Tags			Пользователи
//	@Produce		json
//	@Success		200	{object}	utils.HTTPSuccess
//	@Failure		404	{object}	utils.HTTPError
//	@Param			id	path		string	true	"Идентификатор пользователя"
//	@Router			/api/v1/users/delete/{id} [delete]
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	_, err := user.DeleteUser(id)
	if err != nil {
		utils.NewError(c, http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Пользователь успешно удален"})
	fmt.Println("Пользователь удален")
}
