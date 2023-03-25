package groupCategory

import (
	"event-backend/model/input"
	groupCategoryServices "event-backend/service/groupCategory"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateController(c *gin.Context) {
	var input input.UpdateGroupCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON", "data": err.Error()})
		return
	}

	err := groupCategoryServices.Update(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error while updating groupCategory", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "GroupCategory updated!", "data": ""})
}
