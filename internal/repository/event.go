package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/raphaelmb/go-passin/internal/database/sqlc"
	"github.com/raphaelmb/go-passin/internal/entity"
)

type EventRepository interface {
	CreateEvent(ctx context.Context, e *entity.Event) error
	GetEventByID(ctx context.Context, id uuid.UUID) (*entity.Event, error)
	GetEventBySlug(ctx context.Context, slug string) (*entity.Event, error)
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

func (r *repository) GetEventByID(ctx context.Context, id uuid.UUID) (*entity.Event, error) {
	event, err := r.queries.GetEventByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &entity.Event{
		ID:               event.ID.String(),
		Title:            event.Title,
		Details:          event.Details,
		Slug:             event.Slug,
		MaximumAttendees: int(event.MaximumAttendees),
		CreatedAt:        event.CreatedAt,
		UpdatedAt:        event.UpdatedAt,
	}, nil
}

func (r *repository) GetEventBySlug(ctx context.Context, slug string) (*entity.Event, error) {
	event, err := r.queries.GetEventBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}
	return &entity.Event{
		ID:               event.ID.String(),
		Title:            event.Title,
		Details:          event.Details,
		Slug:             event.Slug,
		MaximumAttendees: int(event.MaximumAttendees),
		CreatedAt:        event.CreatedAt,
		UpdatedAt:        event.UpdatedAt,
	}, nil
}
