package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time-tracker/model"
)

type userResponse struct {
	Data []model.User `json:"data"`
}

type errResponse struct {
	Message string
}

var user model.User

// CreateUser Создает нового пользователя
// @Summary Создает нового пользователя
// @Description Создает нового пользователя
// @Tags Users
// @Accept json
// @Produce json
// @Param user body model.User true "Новый пользователь"
// @Success 201 {object} userResponse
// @Failure 400 {object} errResponse
// @Router /users/create [post]
func CreateUser(c *gin.Context) {
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedUser, err := user.Save()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": savedUser})
	log.Println("Пользователь был удачно создан")
}

// ChangeUserData Изменение данных пользователя
// @Summary Изменение данных пользователя
// @Description Изменение данных пользователя
// @Tags Users
// @Accept json
// @Produce json
// @Router /users/change [patch]
func ChangeUserData(c *gin.Context) {

}

// GetUsersList Получение данных о всех пользователях
// @Summary Получение данных о всех пользователях
// @Description Получение данных о всех пользователях
// @Tags Users
// @Accept json
// @Produce json
// @Router /users/list [get]
func GetUsersList(c *gin.Context) {

}

// GetUserById Получение данных о пользователе
// @Summary Получение данных о пользователе по ID
// @Description Получение данных о пользователе по ID
// @Tags Users
// @Accept json
// @Produce json
// @Router /users/list/{id} [get]
func GetUserById(c *gin.Context) {

}

// DeleteUser Удаление пользователя
// @Summary Удаление пользователя
// @Description Изменение данных пользователя
// @Tags Users
// @Accept json
// @Produce json
// @Router /users/delete [delete]
func DeleteUser(c *gin.Context) {

}
