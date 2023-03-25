package project

import (
	"gradpanel-backend/models/inputs"
	projectServices "gradpanel-backend/services/project"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateController(c *gin.Context) {

	var input inputs.UpdateProjectInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing data", "data": err.Error()})
		return
	}

	err := projectServices.Update(&input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal error", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project updated!", "data": ""})

}
