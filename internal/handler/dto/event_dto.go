package dto

type EventDTO struct {
	Title            string `json:"title" validate:"required"`
	Details          string `json:"details" validate:"required"`
	MaximumAttendees int    `json:"maximumAttendees" validate:"required"`
}

type EventResponseDTO struct {
	ID               string `json:"id"`
	Title            string `json:"title"`
	Details          string `json:"details"`
	Slug             string `json:"slug"`
	MaximumAttendees int    `json:"maximumAttendees"`
	AttendeesAmount  int    `json:"attendeesAmount"`
}

type RegisterForEventDTO struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}
