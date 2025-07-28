package handlers

import (
    "net/http"
    "time"
	"fmt"
    "github.com/gin-gonic/gin"
    "vigovia/models"
    "vigovia/utils"
)

func GenerateItinerary(c *gin.Context) {
    var req models.ItineraryData
	fmt.Println("Received request to generate itinerary",req)
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    filename := "itinerary_" + time.Now().Format("20060102150405") + ".pdf"
    filepath := "storage/pdf/" + filename

    err := utils.GenerateStyledPDF(req, filepath)
    if err != nil {
        fmt.Println("Error generating PDF:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate PDF"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "pdf_url": "/" + filepath,
    })
}
