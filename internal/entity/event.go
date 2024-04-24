package entity

import "time"

type Event struct {
	ID               string    `json:"id"`
	Title            string    `json:"title"`
	Details          string    `json:"details"`
	Slug             string    `json:"slug"`
	MaximumAttendees int       `json:"maximum_attendees"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
