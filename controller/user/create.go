package user

import (
	"gradpanel-backend/models"
	"gradpanel-backend/models/inputs"
	userServices "gradpanel-backend/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateController(c *gin.Context) {
	var input inputs.UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing fields!", "data": err.Error()})
		return
	}

	var user models.User
	user.Mail = input.Mail
	user.Password = input.Password
	user.Role = input.Role

	result, err := userServices.Create(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can not create user", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created!", "data": result})
}
