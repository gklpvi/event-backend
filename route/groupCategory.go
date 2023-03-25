package route

import (
	"event-backend/controller/groupCategory"

	"github.com/gin-gonic/gin"
)

func GroupCategoryRouter(r *gin.RouterGroup) {
	routes := r.Group("/groupCategory")
	{
		routes.POST("/", groupCategory.CreateController)
		routes.PUT("/", groupCategory.UpdateController)
		routes.GET("/", groupCategory.GetController)
		routes.GET("/all", groupCategory.GetAllController)
		routes.DELETE("/", groupCategory.DeleteController)
	}
}
