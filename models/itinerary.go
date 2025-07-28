package models

type TripDetails struct {
    TravelerName     string `json:"travelerName" binding:"required"`
    DepartureCity    string `json:"departureCity" binding:"required"`
    DestinationCity  string `json:"destinationCity" binding:"required"`
    DepartureDate    string `json:"departureDate" binding:"required"`
    ArrivalDate      string `json:"arrivalDate" binding:"required"`
    NumberOfTravelers string `json:"numberOfTravelers" binding:"required"`
}

type DailyItinerary struct {
    ID                string `json:"id" binding:"required"`
    DayTitle          string `json:"dayTitle" binding:"required"`
    Date              string `json:"date" binding:"required"`
    MorningActivities string `json:"morningActivities"`
    AfternoonActivities string `json:"afternoonActivities"`
    EveningActivities string `json:"eveningActivities"`
}

type Flight struct {
    ID      string `json:"id" binding:"required"`
    Date    string `json:"date" binding:"required"`
    Airline string `json:"airline" binding:"required"`
    From    string `json:"from" binding:"required"`
    To      string `json:"to" binding:"required"`
}

type Accommodation struct {
    ID             string `json:"id" binding:"required"`
    City           string `json:"city" binding:"required"`
    HotelName      string `json:"hotelName" binding:"required"`
    CheckInDate    string `json:"checkInDate" binding:"required"`
    CheckOutDate   string `json:"checkOutDate" binding:"required"`
    NumberOfNights int    `json:"numberOfNights" binding:"required"`
}

type Activity struct {
    ID          string `json:"id" binding:"required"`
    City        string `json:"city" binding:"required"`
    ActivityName string `json:"activityName" binding:"required"`
    DateTime    string `json:"dateTime" binding:"required"`
    Duration    string `json:"duration" binding:"required"`
}

type Installment struct {
    ID      string `json:"id" binding:"required"`
    Amount  string `json:"amount" binding:"required"`
    DueDate string `json:"dueDate" binding:"required"`
}

type VisaDetails struct {
    VisaType       string `json:"visaType" binding:"required"`
    ValidityPeriod string `json:"validityPeriod" binding:"required"`
    ProcessingDate string `json:"processingDate" binding:"required"`
}

type PaymentPlan struct {
    TotalAmount string `json:"totalAmount" binding:"required"`
    TCS         string `json:"tcs" binding:"required"`
}

type ItineraryData struct {
    TripDetails      TripDetails      `json:"tripDetails" binding:"required"`
    DailyItineraries []DailyItinerary `json:"dailyItineraries"`
    Flights          []Flight         `json:"flights"`
    Accommodations   []Accommodation  `json:"accommodations"`
    Activities       []Activity       `json:"activities"`
    VisaDetails      VisaDetails      `json:"visaDetails" binding:"required"`
    PaymentPlan      PaymentPlan      `json:"paymentPlan" binding:"required"`
    Installments     []Installment    `json:"installments"`
}
