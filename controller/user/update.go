package user

import (
	"gradpanel-backend/models/inputs"
	userServices "gradpanel-backend/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateController(c *gin.Context) {
	var input inputs.UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing fields!", "data": err.Error()})
		return
	}

	err := userServices.Update(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can not update user", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated!", "data": ""})
}
