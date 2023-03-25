package route

import (
	"gradpanel-backend/models"
	"gradpanel-backend/models/inputs"
	"gradpanel-backend/services/auth"
	"gradpanel-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(c *gin.Context) {
	var input inputs.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON", "data": err.Error()})
		return
	}

	var user models.User
	user.Mail = input.Mail
	user.Role = input.Role

	err := utils.MailValidation(user.Mail)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid mail", "data": err.Error()})
		return
	}

	password, err := utils.GenerateHash(input.Password)
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
