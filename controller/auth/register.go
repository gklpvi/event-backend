package auth

import (
	"event-backend/model"
	"event-backend/service/auth"
	"event-backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Mail     string `json:"mail" binding:"required"`
	Password string `json:"password" binding:"required"`
}
func RegisterController(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON", "data": err.Error()})
		return
	}

	var user model.User
	user.Mail = input.Mail

	valid := util.MailValidation(user.Mail)
	if valid != true {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid mail", "data": ""})
		return
	}

	password, err := util.GenerateHash(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal error!", "data": err.Error()})
		return
	}

	user.Password = password
	u, err := auth.RegisterService(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal error!", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered!", "data": u})
}
