package route

import (
	"gradpanel-backend/controller/profile"

	"github.com/gin-gonic/gin"
)

func ProfileRouter(r *gin.RouterGroup) {
	routes := r.Group("/profile")
	{
		routes.POST("/", profile.CreateController)
		routes.PUT("/", profile.UpdateController)
		routes.GET("/", profile.GetByIdController)
		routes.GET("/by-token", profile.GetByTokenController)
		routes.GET("/all", profile.GetAllController)
		routes.DELETE("/", profile.DeleteController)
	}
}
