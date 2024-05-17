package dto

type AttendeeResponseDTO struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	EventTitle string `json:"eventTitle,omitempty"`
	CheckInURL string `json:"checkInURL,omitempty"`
}
