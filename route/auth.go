package route

import (
	"event-backend/controller/auth"
	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.RouterGroup) {
	routes := r.Group("/user")
	{
		routes.POST("/login", auth.LoginController)
		routes.POST("/register", auth.RegisterController)
	}
}
