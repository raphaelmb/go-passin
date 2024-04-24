package repository

import (
	"context"

	"github.com/raphaelmb/go-passin/internal/database/sqlc"
	"github.com/raphaelmb/go-passin/internal/entity"
)

type EventRepository interface {
	CreateEvent(ctx context.Context, e *entity.Event) error
}

type repository struct {
	queries *sqlc.Queries
}

func NewEventRepository(q *sqlc.Queries) EventRepository {
	return &repository{
		queries: q,
	}
}

func (r *repository) CreateEvent(ctx context.Context, e *entity.Event) error {
	err := r.queries.CreateEvent(ctx, sqlc.CreateEventParams{
		Title:            e.Title,
		Details:          e.Details,
		Slug:             e.Slug,
		MaximumAttendees: int32(e.MaximumAttendees),
	})
	if err != nil {
		return err
	}
	return nil
}
