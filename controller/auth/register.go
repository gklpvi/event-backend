package auth

import (
	"event-backend/model"
	"event-backend/service/auth"
	profileServices "event-backend/service/profile"
	userServices "event-backend/service/user"
	"event-backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Mail     string `json:"mail" binding:"required"`
	Password string `json:"password" binding:"required"`
}
func RegisterController(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON", "data": err.Error()})
		return
	}

	fetchedUser, _ := userServices.GetByMail(input.Mail);
	if fetchedUser.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Mail already exists", "data": ""})
		return
	}
	
	var user model.User
	user.Mail = input.Mail

	valid := util.MailValidation(user.Mail)
	if !valid {
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
	
	profile := model.Profile{
		UserID: u.ID,
		Username: input.Username,
	}
	if _, err = profileServices.Create(&profile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal error!", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered!", "data": u})
}
