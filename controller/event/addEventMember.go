package event

import (
	"event-backend/model"
	"event-backend/service/event"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Input struct {
	EventID  uint   `json:"event_id" binding:"required"`
}
func AddEventMemberController(c *gin.Context) {
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON", "data": err.Error()})
		return
	}
	if err := event.AddEventMemberService(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal error!", "data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event member added successfully", "data": input})
}