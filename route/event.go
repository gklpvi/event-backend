package route

import (
	"event-backend/controller/event"

	"github.com/gin-gonic/gin"
)

func EventRouter(r *gin.RouterGroup) {
	routes := r.Group("/event")
	{
		routes.POST("/", event.CreateController)
		routes.POST("/add-event-member", event.AddEventMemberController)
		routes.PUT("/", event.UpdateController)
		routes.GET("/", event.GetController)
		routes.GET("/all", event.GetAllController)
		routes.DELETE("/", event.DeleteController)
	}
}
