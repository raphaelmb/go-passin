package entity

import "time"

type Event struct {
	ID               string
	Title            string
	Details          string
	Slug             string
	MaximumAttendees int
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
