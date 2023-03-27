package event

import (
	"event-backend/model"
	eventServices "event-backend/service/event"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GroupInput struct {
	GroupCategoryID int    `json:"eventCategoryID" binding:"required"`
	MaxMember       int    `json:"maxMember" binding:"required"`
}
func CreateController(c *gin.Context) {
	var input GroupInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing data", "data": err.Error()})
		return
	}

	var event model.Event
	result, err := eventServices.Create(&event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal error", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event is created!", "data": result})
}
