package project

import (
	projectServices "gradpanel-backend/services/project"
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

	projects, err := projectServices.GetById(uintId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can not get project by id","data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, projects)
}


func GetAllController(c *gin.Context) {

	projects, err := projectServices.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can not get all projects","data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": projects})
}

func GetDetailsController(c *gin.Context) {
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

	detail, err := projectServices.GetDetails(uintId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can not get project details", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "get project details successfully", "data": detail})
}
