package routes

import (
	"Events/api/handlers"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine, eventHandler *handlers.EventHandler) {
    router.POST("/events", eventHandler.AddEvent)
    router.PUT("/events/:id", eventHandler.UpdateEvent)
    router.GET("/events", eventHandler.GetEvents)
    router.DELETE("/events/:id", eventHandler.DeleteEvent)
    router.GET("/events/:id", eventHandler.GetEventById)
    router.PUT("/events/:id/:clasification", eventHandler.ManageEvent)
}

