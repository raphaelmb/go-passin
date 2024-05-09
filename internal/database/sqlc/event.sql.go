// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: event.sql

package sqlc

import (
	"context"

	"github.com/google/uuid"
)

const createEvent = `-- name: CreateEvent :exec
INSERT INTO events(title, details, slug, maximum_attendees) VALUES($1, $2, $3, $4)
`

type CreateEventParams struct {
	Title            string
	Details          string
	Slug             string
	MaximumAttendees int32
}

func (q *Queries) CreateEvent(ctx context.Context, arg CreateEventParams) error {
	_, err := q.db.ExecContext(ctx, createEvent,
		arg.Title,
		arg.Details,
		arg.Slug,
		arg.MaximumAttendees,
	)
	return err
}

const getEventByID = `-- name: GetEventByID :one
SELECT id, title, details, slug, maximum_attendees, created_at, updated_at FROM events WHERE id = $1
`

func (q *Queries) GetEventByID(ctx context.Context, id uuid.UUID) (Event, error) {
	row := q.db.QueryRowContext(ctx, getEventByID, id)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Details,
		&i.Slug,
		&i.MaximumAttendees,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getEventBySlug = `-- name: GetEventBySlug :one
SELECT id, title, details, slug, maximum_attendees, created_at, updated_at FROM events WHERE slug = $1
`

func (q *Queries) GetEventBySlug(ctx context.Context, slug string) (Event, error) {
	row := q.db.QueryRowContext(ctx, getEventBySlug, slug)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Details,
		&i.Slug,
		&i.MaximumAttendees,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
