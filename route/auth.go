package route

import (
	"event-backend/controller/auth"

	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.RouterGroup) {
	r.POST("/register", auth.RegisterController)
	r.POST("/login", auth.LoginController)
}
