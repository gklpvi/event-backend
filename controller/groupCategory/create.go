package groupCategory

import (
	"event-backend/model"
	groupCategoryServices "event-backend/service/groupCategory"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GroupCategoryInput struct {
	Name     string `json:"name" binding:"required"`
	MaxLevel int    `json:"maxLevel" binding:"required"`
}

func CreateController(c *gin.Context) {
	var input GroupCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON", "data": err.Error()})
		return
	}

	var groupCategory model.GroupCategory
	groupCategory.Name = input.Name
	groupCategory.MaxLevel = input.MaxLevel

	result, err := groupCategoryServices.Create(&groupCategory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error while creating groupCategory", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "GroupCategory created!", "data": result})
}
