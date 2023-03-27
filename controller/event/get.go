package event

import (
	eventServices "event-backend/service/event"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetController(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing id!", "data": ""})
		return
	}
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing id!", "data": ""})
		return
	}
	uintId := uint(uid)

	events, err := eventServices.GetById(uintId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can not get event by id", "data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}

func GetAllController(c *gin.Context) {

	events, err := eventServices.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can not get all events", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": events})
}
