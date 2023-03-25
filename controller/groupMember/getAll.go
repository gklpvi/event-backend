package groupMember

import (
	groupMemberServices "event-backend/service/groupMember"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllController(c *gin.Context) {

	groupMembers, err := groupMemberServices.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error while getting groupMembers", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": groupMembers})

}
