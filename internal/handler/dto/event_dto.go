package dto

import "time"

type EventDTO struct {
	Title            string `json:"title" validate:"required"`
	Details          string `json:"details" validate:"required"`
	MaximumAttendees int    `json:"maximum_attendees" validate:"required"`
}

type EventResponseDTO struct {
	ID               string    `json:"id"`
	Title            string    `json:"title"`
	Details          string    `json:"details"`
	Slug             string    `json:"slug"`
	MaximumAttendees int       `json:"maximum_attendees"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
