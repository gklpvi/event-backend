package projectStudent

import (
	projectStudentServices "gradpanel-backend/services/projectStudent"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetController(c *gin.Context) {

	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing id!", "data": ""})
		return
	}
	uid, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing id!", "data": ""})
		return
	}
	uintId := uint(uid)

	projectStudents, err := projectStudentServices.GetById(uintId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can not get projectStudent by id", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": projectStudents})
}
