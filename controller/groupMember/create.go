package projectStudent

import (
	"gradpanel-backend/models"
	"gradpanel-backend/models/inputs"
	projectStudentServices "gradpanel-backend/services/projectStudent"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateController(c *gin.Context) {
	var input inputs.ProjectStudentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON", "data": err.Error()})
		return
	}

	var projectStudent models.ProjectStudent
	projectStudent.ProjectID = input.ProjectID
	projectStudent.ProfileID = input.ProfileID

	result, err := projectStudentServices.Create(&projectStudent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error while creating projectStudent", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ProjectStudent created!", "data": result})
}
