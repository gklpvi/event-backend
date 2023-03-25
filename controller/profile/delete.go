package profile

import (
	profileServices "gradpanel-backend/services/profile"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteController(c *gin.Context) {
	type Input struct {
		Id uint `json:"id" binding:"required"`
	}
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing fields!", "data": err.Error()})
		return
	}

	err := profileServices.Delete(input.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can not delete profile", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile deleted!", "data": ""})
}
