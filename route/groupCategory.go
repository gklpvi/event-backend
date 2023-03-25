package route

import (
	"gradpanel-backend/controller/category"
	
	"github.com/gin-gonic/gin"
)

func CategoryRouter(r *gin.RouterGroup) {
	routes := r.Group("/category")
	{
		routes.POST("/", category.CreateController)
		routes.PUT("/", category.UpdateController)
		routes.GET("/", category.GetController)
		routes.GET("/all", category.GetAllController)
		routes.DELETE("/", category.DeleteController)
	}
}
