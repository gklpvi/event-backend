package status

import (
	"gradpanel-backend/models"
	"gradpanel-backend/models/inputs"
	statusServices "gradpanel-backend/services/status"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateController(c *gin.Context) {
	var input inputs.StatusInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON", "data": err.Error()})
		return
	}

	var status models.Status
	status.Name = input.Name
	status.Type = input.Type

	result, err := statusServices.Create(&status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error while creating status", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status created!", "data": result})
}
