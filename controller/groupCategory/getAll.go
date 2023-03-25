package groupCategory

import (
	groupCategoryServices "event-backend/service/groupCategory"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllController(c *gin.Context) {
	groupCategory, err := groupCategoryServices.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error while getting groupCategory", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": groupCategory})

}
