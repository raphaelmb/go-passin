package response

import "time"

type EventResponse struct {
	ID               string    `json:"id"`
	Title            string    `json:"title" validate:"required"`
	Details          string    `json:"details"`
	Slug             string    `json:"slug"`
	MaximumAttendees int       `json:"maximum_attendees" validate:"required"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
