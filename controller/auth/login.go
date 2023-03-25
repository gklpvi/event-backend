package auth

import (
	"event-backend/service/auth"
	"event-backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Mail     string `json:"mail" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginController(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON", "data": err.Error()})
		return
	}

	user, err := auth.LoginService(input.Mail)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found!", "data": err.Error()})
		return
	}

	if err := util.VerifyPassword(input.Password, user.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Wrong mail or password", "data": err.Error()})
		return
	}

	token, err := util.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal error!", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "data": user, "token": token})
}
