package projectStudent

import (
	applicationServices "gradpanel-backend/services/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllController(c *gin.Context) {

	applications, err := applicationServices.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error while getting applications", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": applications})

}
