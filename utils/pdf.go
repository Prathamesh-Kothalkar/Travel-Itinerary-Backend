package utils

import (
    "fmt"
    "vigovia/models"

    "github.com/jung-kurt/gofpdf"
)

func GeneratePDF(data models.ItineraryRequest, filepath string) error {
    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()
    pdf.SetFont("Arial", "B", 16)
    pdf.Cell(40, 10, "Travel Itinerary")

    pdf.Ln(12)
    pdf.SetFont("Arial", "", 12)
    pdf.Cell(0, 10, fmt.Sprintf("Name: %s", data.User.Name))
    pdf.Ln(8)
    pdf.Cell(0, 10, fmt.Sprintf("Hotel: %s, %s", data.Hotel.Name, data.Hotel.Location))
    pdf.Ln(8)
    pdf.Cell(0, 10, fmt.Sprintf("Check-In: %s", data.Hotel.CheckIn))
    pdf.Ln(8)
    pdf.Cell(0, 10, fmt.Sprintf("Check-Out: %s", data.Hotel.CheckOut))
    pdf.Ln(8)

    pdf.Cell(0, 10, "Activities:")
    for _, act := range data.Activities {
        pdf.Ln(6)
        pdf.Cell(0, 10, "- "+act)
    }

    return pdf.OutputFileAndClose(filepath)
}
