package project

import (
	projectServices "gradpanel-backend/services/project"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteController(c *gin.Context) {
	type Input struct {
		Id uint `json:"id" binding:"required"`
	}
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing data", "data": err.Error()})
		return
	}

	err := projectServices.Delete(input.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal error", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project deleted!", "data": ""})
}
