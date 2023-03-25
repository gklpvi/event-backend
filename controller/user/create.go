package user

import (
	userServices "event-backend/service/user"
	"event-backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInput struct {
	Mail     string `json:"mail" binding:"required"`
	Password string `json:"password" binding:"required"`
}
func CreateController(c *gin.Context) {
	var input UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing fields!", "data": err.Error()})
		return
	}

	var user model.User
	user.Mail = input.Mail
	user.Password = input.Password

	result, err := userServices.Create(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can not create user", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created!", "data": result})
}
