package projectStudent

import (
	projectStudentServices "gradpanel-backend/services/projectStudent"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteController(c *gin.Context) {
	type Input struct {
		Id uint `json:"id"`
	}
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON", "data": err.Error()})
		return
	}

	err := projectStudentServices.Delete(input.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error while deleting projectStudent", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ProjectStudent deleted!", "data": ""})
}
