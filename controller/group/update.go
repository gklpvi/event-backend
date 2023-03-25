package group

import (
	groupCategoryServices "event-backend/service/groupCategory"
	"event-backend/model/input"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateController(c *gin.Context) {

	var input input.UpdateGroupCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing data", "data": err.Error()})
		return
	}

	err := groupCategoryServices.Update(&input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal error", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Group updated!", "data": ""})

}
