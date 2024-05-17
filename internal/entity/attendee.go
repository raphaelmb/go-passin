package entity

import "time"

type Attendee struct {
	ID         int
	Name       string
	Email      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	EventTitle string
}
