package status

import (
	statusServices "gradpanel-backend/services/status"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteController(c *gin.Context) {
	type Input struct {
		Id uint `json:"id" binding:"required"`
	}
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON", "data": err.Error()})
		return
	}

	err := statusServices.Delete(input.Id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error while deleting status", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status deleted!", "data": ""})
}
