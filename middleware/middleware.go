package middleware

import (
	"event-backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BlacklistTokenMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){		
	tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(401, gin.H{"message": "Token is missing"})
			return
		}
		if util.TokenBlacklist[tokenString] {
			c.AbortWithStatusJSON(401, gin.H{"message": "Token is blacklisted"})
			return
		}
		c.Next()
	}
}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := util.TokenValid(c)
		if err != nil {
			tokenString := util.ExtractToken(c)
			util.RemoveFromWhitelist(tokenString)
			util.AddToBlacklist(tokenString)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized", "data": err.Error(),})
			return
		}
		c.Next()
	}
}
