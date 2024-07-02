package handlers

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
