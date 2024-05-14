package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/raphaelmb/go-passin/internal/database/sqlc"
	"github.com/raphaelmb/go-passin/internal/entity"
	"github.com/raphaelmb/go-passin/internal/handler/dto"
)

type EventRepository interface {
	CreateEvent(ctx context.Context, e *entity.Event) error
	GetEventByID(ctx context.Context, id uuid.UUID) (*entity.Event, error)
	GetEventBySlug(ctx context.Context, slug string) (*entity.Event, error)
	RegisterForEvent(ctx context.Context, e dto.RegisterForEventDTO, id uuid.UUID) error
}

type EventRepo struct {
	queries *sqlc.Queries
}

func NewEventRepository(q *sqlc.Queries) EventRepository {
	return &EventRepo{
		queries: q,
	}
}

func (r *EventRepo) CreateEvent(ctx context.Context, e *entity.Event) error {
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

func (r *EventRepo) GetEventByID(ctx context.Context, id uuid.UUID) (*entity.Event, error) {
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

func (r *EventRepo) GetEventBySlug(ctx context.Context, slug string) (*entity.Event, error) {
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

func (r *EventRepo) RegisterForEvent(ctx context.Context, e dto.RegisterForEventDTO, id uuid.UUID) error {
	err := r.queries.RegisterForEvent(ctx, sqlc.RegisterForEventParams{
		Name:    e.Name,
		Email:   e.Email,
		EventID: id,
	})
	if err != nil {
		return err
	}
	return nil
}
