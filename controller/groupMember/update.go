package groupMember

import (
	"event-backend/model/input"
	groupMemberServices "event-backend/service/groupMember"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateController(c *gin.Context) {
	var input input.UpdateGroupMemberInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON", "data": err.Error()})
		return
	}

	err := groupMemberServices.Update(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error while updating groupMember", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "GroupMember updated!", "data": ""})
}
