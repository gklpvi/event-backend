package profile

import (
	"gradpanel-backend/models/inputs"
	profileServices "gradpanel-backend/services/profile"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateController(c *gin.Context) {
	var input inputs.UpdateProfileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing fields!", "data": err.Error()})
		return
	}

	err := profileServices.Update(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can not update profile", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated!", "data": ""})
}
