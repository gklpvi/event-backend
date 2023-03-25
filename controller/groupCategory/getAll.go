package status

import (
	statusServices "gradpanel-backend/services/status"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllController(c *gin.Context) {
	status, err := statusServices.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error while getting status", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": status})

}
