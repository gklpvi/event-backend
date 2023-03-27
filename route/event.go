package route

import (
	"event-backend/controller/event"

	"github.com/gin-gonic/gin"
)

func EventRouter(r *gin.RouterGroup) {
	routes := r.Group("/event")
	{
		routes.POST("/add-event-member", event.AddEventMemberController)
	}
}
