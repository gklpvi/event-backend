package status

import (
	"gradpanel-backend/models/inputs"
	statusServices "gradpanel-backend/services/status"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateController(c *gin.Context) {
	var input inputs.UpdateStatusInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON", "data": err.Error()})
		return
	}

	err := statusServices.Update(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error while updating status", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status updated!", "data": ""})
}
