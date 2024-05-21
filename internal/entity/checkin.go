package entity

import "time"

type CheckIn struct {
	ID         int
	CreatedAt  time.Time
	AttendeeID int
}
