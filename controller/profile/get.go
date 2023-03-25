package profile

import (
	profileServices "event-backend/service/profile"
	"event-backend/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetByIdController(c *gin.Context) {
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

	profiles, err := profileServices.GetById(uintId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can not get profile", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": profiles})
}

func GetByTokenController(c *gin.Context) {
	profileId, err := util.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "can not get profile id", "data": err.Error()})
		return
	}

	profile, err := profileServices.GetById(profileId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "can not get profile", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": profile})
}

func GetAllController(c *gin.Context) {
	profiles, err := profileServices.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can not get all profiles", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": profiles})
}
