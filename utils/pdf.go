package utils

import (
	"fmt"
	"vigovia/models"
    "time"
	"github.com/jung-kurt/gofpdf"
)

func GenerateStyledPDF(data models.ItineraryData, filepath string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// --- HEADER: Logo + Banner ---
	logoPath := "public/images/logo.png" // Place logo in assets folder
	pdf.ImageOptions(logoPath, 80, 10, 50, 20, false, gofpdf.ImageOptions{}, 0, "")
	pdf.Ln(25)

	// Banner background
	pdf.SetFillColor(147, 111, 224) // Purple banner
	pdf.Rect(10, 35, 190, 25, "F")
	pdf.SetTextColor(255, 255, 255)
	pdf.SetFont("Arial", "B", 16)
	pdf.SetXY(10, 40)
	pdf.CellFormat(190, 10, fmt.Sprintf("Hi %s! Your %s Itinerary",
		data.TripDetails.TravelerName, data.TripDetails.DestinationCity), "", 1, "C", false, 0, "")
	pdf.Ln(10)

     // TRIP SUMMARY BAR
    pdf.SetY(70)
    pdf.SetFillColor(255, 255, 255)
    pdf.SetDrawColor(138, 43, 226)
    pdf.SetLineWidth(0.5)
    pdf.Rect(15, pdf.GetY(), 180, 25, "D")
    pdf.SetFont("Helvetica", "B", 12)
    pdf.SetTextColor(0, 0, 0)
    pdf.SetXY(20, pdf.GetY()+5)
    pdf.CellFormat(40, 8, "Departure From", "", 0, "L", false, 0, "")
    pdf.CellFormat(40, 8, "Departure Date", "", 0, "L", false, 0, "")
    pdf.CellFormat(40, 8, "Arrival Date", "", 0, "L", false, 0, "")
    pdf.CellFormat(40, 8, "Travellers", "", 1, "L", false, 0, "")

    pdf.SetFont("Helvetica", "", 12)
    pdf.SetX(20)
    pdf.CellFormat(40, 8, data.TripDetails.DepartureCity, "", 0, "L", false, 0, "")
    pdf.CellFormat(40, 8, formatDate(data.TripDetails.DepartureDate), "", 0, "L", false, 0, "")
    pdf.CellFormat(40, 8, formatDate(data.TripDetails.ArrivalDate), "", 0, "L", false, 0, "")

	// --- Trip Summary Bar ---
	pdf.SetFillColor(255, 255, 255)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Arial", "", 11)
	pdf.SetDrawColor(138, 43, 226)
	pdf.Rect(10, 65, 190, 20, "D")

	pdf.SetXY(15, 70)
	pdf.Cell(50, 5, fmt.Sprintf("Departure From: %s", data.TripDetails.DepartureCity))
	pdf.SetXY(80, 70)
	pdf.Cell(50, 5, fmt.Sprintf("Departure: %s", data.TripDetails.DepartureDate))
	pdf.SetXY(150, 70)
	pdf.Cell(50, 5, fmt.Sprintf("Travelers: %s", data.TripDetails.NumberOfTravelers))
	pdf.Ln(25)

	// --- Accommodations Table ---
	if len(data.Accommodations) > 0 {
		addTableSection(pdf, "Hotel Bookings", []string{"City", "Hotel", "Check-In", "Check-Out", "Nights"})
		for _, acc := range data.Accommodations {
			addTableRow(pdf, []string{
				acc.City, acc.HotelName, acc.CheckInDate, acc.CheckOutDate, fmt.Sprintf("%d", acc.NumberOfNights),
			})
		}
		pdf.Ln(10)
	}

	// --- Activities Table ---
	if len(data.Activities) > 0 {
		addTableSection(pdf, "Activities", []string{"City", "Activity", "Date/Time", "Duration"})
		for _, act := range data.Activities {
			addTableRow(pdf, []string{
				act.City, act.ActivityName, act.DateTime, act.Duration,
			})
		}
		pdf.Ln(10)
	}
    //fligths
    if len(data.Flights) > 0 {
        pdf.Ln(15)
        pdf.SetFont("Helvetica", "B", 14)
        pdf.SetTextColor(0, 0, 0)
        pdf.Cell(0, 10, "Flight Summary")

        for _, flight := range data.Flights {
            pdf.Ln(8)
            pdf.SetDrawColor(195, 155, 211)
            pdf.Rect(15, pdf.GetY(), 180, 12, "D")
            pdf.SetXY(20, pdf.GetY()+3)
            pdf.SetFont("Helvetica", "", 11)
            pdf.CellFormat(0, 6, fmt.Sprintf("%s: Fly %s from %s to %s", formatDayMonth(flight.Date), flight.Airline, flight.From, flight.To), "", 0, "L", false, 0, "")
        }
    }

	// --- Visa Details ---
	pdf.SetFillColor(91, 44, 111)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetFont("Arial", "B", 14)
	pdf.CellFormat(190, 10, "Visa Details", "1", 1, "C", true, 0, "")
	pdf.SetFont("Arial", "", 12)
	pdf.SetTextColor(0, 0, 0)
	pdf.Cell(0, 8, fmt.Sprintf("Visa Type: %s", data.VisaDetails.VisaType))
	pdf.Ln(6)
	pdf.Cell(0, 8, fmt.Sprintf("Validity Period: %s", data.VisaDetails.ValidityPeriod))
	pdf.Ln(6)
	pdf.Cell(0, 8, fmt.Sprintf("Processing Date: %s", data.VisaDetails.ProcessingDate))
	pdf.Ln(12)

	// --- Payment Plan ---
	pdf.SetFillColor(91, 44, 111)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetFont("Arial", "B", 14)
	pdf.CellFormat(190, 10, "Payment Plan", "1", 1, "C", true, 0, "")
	pdf.SetFont("Arial", "", 12)
	pdf.SetTextColor(0, 0, 0)
	pdf.Cell(0, 8, fmt.Sprintf("Total Amount: ₹%s", data.PaymentPlan.TotalAmount))
	pdf.Ln(6)
	pdf.Cell(0, 8, fmt.Sprintf("TCS: %s", data.PaymentPlan.TCS))
	pdf.Ln(10)

	// --- Installments Table ---
	if len(data.Installments) > 0 {
		addTableSection(pdf, "Installments", []string{"Installment", "Amount", "Due Date"})
		for i, inst := range data.Installments {
			addTableRow(pdf, []string{
				fmt.Sprintf("Installment %d", i+1), "₹" + inst.Amount, inst.DueDate,
			})
		}
		pdf.Ln(10)
	}

	// --- Footer ---
	pdf.SetY(-40)
	pdf.SetFont("Arial", "", 9)
	pdf.SetDrawColor(200, 200, 200)
	pdf.Line(10, pdf.GetY(), 200, pdf.GetY())
	pdf.Ln(5)
	pdf.Cell(0, 5, "Vigovia Tech Pvt. Ltd. | PLAN.PACK.GO!")
	pdf.Ln(4)
	pdf.Cell(0, 5, "Registered Office: Hd-109 Cinnabar Hills, Links Business Park, Karnataka, India.")
	pdf.Ln(4)
	pdf.Cell(0, 5, "Phone: +91-99X9999999 | Email: Contact@vigovia.com")

	return pdf.OutputFileAndClose(filepath)
}

// --- Helper to add table header ---
func addTableSection(pdf *gofpdf.Fpdf, title string, headers []string) {
	pdf.SetFillColor(91, 44, 111)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetFont("Arial", "B", 14)
	pdf.CellFormat(190, 10, title, "1", 1, "C", true, 0, "")

	pdf.SetFont("Arial", "B", 10)
	pdf.SetFillColor(230, 230, 250)
	pdf.SetTextColor(0, 0, 0)
	colWidth := 190.0 / float64(len(headers))
	for _, h := range headers {
		pdf.CellFormat(colWidth, 8, h, "1", 0, "C", true, 0, "")
	}
	pdf.Ln(-1)
}

// --- Helper to add table row ---
func addTableRow(pdf *gofpdf.Fpdf, cols []string) {
	colWidth := 190.0 / float64(len(cols))
	pdf.SetFont("Arial", "", 10)
	for _, c := range cols {
		pdf.CellFormat(colWidth, 8, c, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)
}

func formatDate(dateStr string) string {
    t, err := time.Parse("2006-01-02", dateStr)
    if err != nil {
        return "TBD"
    }
    return t.Format("02 Jan 2006")
}

func formatDayMonth(dateStr string) string {
    t, err := time.Parse("2006-01-02", dateStr)
    if err != nil {
        return "TBD"
    }
    return t.Format("02 Jan")
}
