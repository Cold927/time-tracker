package controller

import (
	"github.com/gin-gonic/gin"
	"log"
)

type errResponse struct {
	Message string
}

// CreateUser Создает нового пользователя
// @Summary Создает нового пользователя
// @Description Создает нового пользователя
// @Security bearerToken
// @Tags Users
// @Accept json
// @Produce json
// @Router /users/create [post]
func CreateUser(c *gin.Context) {
	log.Println("Create User")
}

// ChangeUserData Изменение данных пользователя
// @Summary Изменение данных пользователя
// @Description Изменение данных пользователя
// @Security bearerToken
// @Tags Users
// @Accept json
// @Produce json
// @Router /users/change [patch]
func ChangeUserData(c *gin.Context) {

}

// GetUsersList Получение данных о всех пользователях
// @Summary Получение данных о всех пользователях
// @Description Получение данных о всех пользователях
// @Security bearerToken
// @Tags Users
// @Accept json
// @Produce json
// @Router /users/list [get]
func GetUsersList(c *gin.Context) {

}

// GetUserById Получение данных о пользователе
// @Summary Получение данных о пользователе по ID
// @Description Получение данных о пользователе по ID
// @Security bearerToken
// @Tags Users
// @Accept json
// @Produce json
// @Router /users/list/:id [get]
func GetUserById(c *gin.Context) {

}

// DeleteUser Удаление пользователя
// @Summary Удаление пользователя
// @Description Изменение данных пользователя
// @Security bearerToken
// @Tags Users
// @Accept json
// @Produce json
// @Router /users/delete [delete]
func DeleteUser(c *gin.Context) {

}
