package profile

import (
	"gradpanel-backend/models"
	"gradpanel-backend/models/inputs"
	profileServices "gradpanel-backend/services/profile"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateController(c *gin.Context) {
	var input inputs.ProfileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing fields!", "data": err.Error()})
		return
	}

	var profile models.Profile
	profile.FirstName = input.FirstName
	profile.LastName = input.LastName
	profile.Department = input.Department
	profile.Major = input.Major
	profile.Grade = input.Grade
	profile.Number = input.Number
	profile.GPA = input.GPA
	profile.Biography = input.Biography
	profile.ProfileImage = input.ProfileImage
	profile.UserID = input.UserID

	result, err := profileServices.Create(&profile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can not create profile", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile created!", "data": result})
}
