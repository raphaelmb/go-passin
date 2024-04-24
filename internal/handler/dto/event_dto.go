package dto

type EventDTO struct {
	Title            string `json:"title" validate:"required"`
	Details          string `json:"details"`
	Slug             string `json:"slug"`
	MaximumAttendees int    `json:"maximum_attendees" validate:"required"`
}
