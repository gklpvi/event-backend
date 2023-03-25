package route

import (
	"event-backend/controller/group"

	"github.com/gin-gonic/gin"
)

func GroupRouter(r *gin.RouterGroup) {
	routes := r.Group("/group")
	{
		routes.POST("/", group.CreateController)
		routes.PUT("/", group.UpdateController)
		routes.GET("/", group.GetController)
		routes.GET("/all", group.GetAllController)
		routes.DELETE("/", group.DeleteController)
	}
}
