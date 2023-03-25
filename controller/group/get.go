package group

import (
	groupServices "event-backend/service/group"
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

	groups, err := groupServices.GetById(uintId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can not get group by id", "data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, groups)
}

func GetAllController(c *gin.Context) {

	groups, err := groupServices.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can not get all groups", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": groups})
}
