package project

import (
	"gradpanel-backend/models"
	"gradpanel-backend/models/inputs"
	profileServices "gradpanel-backend/services/profile"
	projectServices "gradpanel-backend/services/project"
	"gradpanel-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateController(c *gin.Context) {
	var input inputs.ProjectInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing data", "data": err.Error()})
		return
	}

	profileId, err := utils.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "extract error", "data": err.Error()})
		return
	}
	instructorProfile, err := profileServices.GetById(profileId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "get by user id for profile error", "data": err.Error()})
		return
	}

	var project models.Project
	project.Title = input.Title
	project.Major = input.Major
	project.Capacity = input.Capacity
	project.Status = input.Status
	project.ProfileID = instructorProfile.ID

	result, err := projectServices.Create(&project)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal error", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project is created!", "data": result})
}
