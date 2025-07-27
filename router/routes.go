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
		api.GET("/itinerary", func(c *gin.Context) { 
			c.JSON(200, gin.H{"message": "Itinerary endpoint is working"})
		})
    }
    return r
}
