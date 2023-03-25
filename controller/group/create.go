package group

import (
	"event-backend/model"
	groupServices "event-backend/service/group"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GroupInput struct {
	GroupCategoryID int    `json:"groupCategoryID" binding:"required"`
	MaxMember       int    `json:"maxMember" binding:"required"`
}
func CreateController(c *gin.Context) {
	var input GroupInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing data", "data": err.Error()})
		return
	}

	var group model.Group
	result, err := groupServices.Create(&group)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal error", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Group is created!", "data": result})
}
