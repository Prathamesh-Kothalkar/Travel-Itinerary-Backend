package models

type User struct {
    Name  string `json:"name" binding:"required"`
    Email string `json:"email" binding:"required,email"`
}

type Hotel struct {
    Name     string `json:"name" binding:"required"`
    Location string `json:"location" binding:"required"`
    CheckIn  string `json:"check_in" binding:"required"`
    CheckOut string `json:"check_out" binding:"required"`
}

type ItineraryRequest struct {
    User       User     `json:"user" binding:"required"`
    Hotel      Hotel    `json:"hotel" binding:"required"`
    Activities []string `json:"activities"`
}
