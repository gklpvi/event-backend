package profile

import (
	"event-backend/model"
	profileServices "event-backend/service/profile"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileInput struct {
	Username string `json:"username" binding:"required"`
	UserID   uint   `json:"userID" binding:"required"`
	Level    int    `json:"level" binding:"required"`
}

func CreateController(c *gin.Context) {
	var input ProfileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing fields!", "data": err.Error()})
		return
	}

	var profile model.Profile
	profile.Username = input.Username
	profile.UserID = input.UserID
	profile.Level = input.Level

	result, err := profileServices.Create(&profile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can not create profile", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile created!", "data": result})
}
