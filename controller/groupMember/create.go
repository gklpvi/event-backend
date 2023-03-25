package groupMember

import (
	"event-backend/model"
	groupMemberServices "event-backend/service/groupMember"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GroupMemberInput struct {
	GroupID  uint `json:"groupId"`
	ProfileID uint `json:"profileId"`
}
func CreateController(c *gin.Context) {
	var input GroupMemberInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON", "data": err.Error()})
		return
	}

	var groupMember model.GroupMember
	groupMember.GroupID = input.GroupID
	groupMember.ProfileID = input.ProfileID

	result, err := groupMemberServices.Create(&groupMember)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error while creating groupMember", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "GroupMember created!", "data": result})
}
