package route

import (
	"gradpanel-backend/controller/project"

	"github.com/gin-gonic/gin"
)

func ProjectRouter(r *gin.RouterGroup) {
	routes := r.Group("/project")
	{
		routes.POST("/", project.CreateController)
		routes.PUT("/", project.UpdateController)
		routes.GET("/", project.GetController)
		routes.GET("/all", project.GetAllController)
		routes.DELETE("/", project.DeleteController)
		routes.GET("/details", project.GetDetailsController)
	}
}
