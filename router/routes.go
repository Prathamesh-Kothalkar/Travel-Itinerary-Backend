package router

import (
    "github.com/gin-gonic/gin"
    "vigovia/handlers"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    api := r.Group("/api/v1")
    {
        api.POST("/itinerary", handlers.GenerateItinerary)
    }
    return r
}
