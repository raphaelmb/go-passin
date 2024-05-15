package dto

type AttendeeResponseDTO struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	EventTitle string `json:"eventTitle,omitempty"`
}
