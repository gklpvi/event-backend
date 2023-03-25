package route

import (
	"event-backend/controller/groupMember"

	"github.com/gin-gonic/gin"
)

func GroupMemberRouter(r *gin.RouterGroup) {
	routes := r.Group("/groupMember")
	{
		routes.POST("/", groupMember.CreateController)
		routes.PUT("/", groupMember.UpdateController)
		routes.GET("/", groupMember.GetController)
		routes.GET("/all", groupMember.GetAllController)
		routes.DELETE("/", groupMember.DeleteController)
	}
}
