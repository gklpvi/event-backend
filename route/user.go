package route

import (
	"gradpanel-backend/controller/user"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	routes := r.Group("/user")
	{
		routes.POST("/", user.CreateController)
		routes.PUT("/", user.UpdateController)
		routes.GET("/", user.GetByIdController)
		routes.GET("/by-token", user.GetByTokenController)
		routes.GET("/all", user.GetAllController)
		routes.DELETE("/", user.DeleteController)
	}
}
