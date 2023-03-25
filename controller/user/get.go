package user

import (
	userServices "gradpanel-backend/services/user"
	profileServices "gradpanel-backend/services/profile"
	"gradpanel-backend/utils"
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

	user, err := userServices.GetById(uintId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "no such user", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user found", "data": user})
}

func GetByTokenController(c *gin.Context) {
	profileId, err := utils.ExtractTokenID(c)	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "can not extract user id", "data": err.Error()})
		return
	}
	
	profile, err := profileServices.GetById(profileId)
	user := profile.User
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "can not get user bu id", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"success","data":user})
}

func GetAllController(c *gin.Context) {
	users, err := userServices.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can not get the data", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": users})
}
