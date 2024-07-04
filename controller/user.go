package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time-tracker/model"
	"time-tracker/utils"
)

type userResponse struct {
	Data []model.User `json:"data"`
}

var user model.User

// CreateUser Создает нового пользователя
//
//	@Summary		Создает нового пользователя
//	@Description	Создает нового пользователя
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		model.User	true	"Новый пользователь"
//	@Success		201		{object}	userResponse
//	@Failure		400		{object}	utils.HTTPError
//	@Router			/users/create [post]
func CreateUser(c *gin.Context) {
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	savedUser, err := user.Save()

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
//	@Param			id		path		string		true	"Идентификатор пользователя"
//	@Param			user	body		model.User	true	"Изменение данных пользователя"
//	@Failure		400		{object}	utils.HTTPError
//	@Router			/users/update/{id} [patch]
func UpdateUserData(c *gin.Context) {
	id := c.Param("id")
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}
	updatedUser, err := user.UpdateData(id, user)
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
//	@Produce		json
//	@Failure		400	{object}	utils.HTTPError
//	@Router			/users/info [get]
func GetUsersInfo(c *gin.Context) {
	fmt.Println("Получение данных о всех пользователях")
}

// GetUserById Получение данных о пользователе
//
//	@Summary		Получение данных о пользователе по ID
//	@Description	Получение данных о пользователе по ID
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		uint	true	"Идентификатор пользователя"
//	@Failure		400	{object}	utils.HTTPError
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
//	@Failure		400	{object}	utils.HTTPError
//	@Param			id	path		uint	true	"Идентификатор пользователя"
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
