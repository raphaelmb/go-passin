package service

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/raphaelmb/go-passin/internal/entity"
	"github.com/raphaelmb/go-passin/internal/handler/dto"
	"github.com/raphaelmb/go-passin/internal/repository"
	"github.com/raphaelmb/go-passin/internal/response"
)

type EventService interface {
	CreateEvent(ctx context.Context, e dto.EventDTO) error
	GetEventByID(ctx context.Context, id uuid.UUID) (*response.EventResponse, error)
}

type service struct {
	repo repository.EventRepository
}

func NewEventService(repo repository.EventRepository) EventService {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateEvent(ctx context.Context, e dto.EventDTO) error {
	err := s.repo.CreateEvent(ctx, &entity.Event{
		Title:            e.Title,
		Details:          e.Details,
		Slug:             e.Slug,
		MaximumAttendees: e.MaximumAttendees,
	})
	if err != nil {
		slog.Error("error creating event", "err", err)
		return err
	}

	return nil
}

func (s *service) GetEventByID(ctx context.Context, id uuid.UUID) (*response.EventResponse, error) {
	event, err := s.repo.GetEventByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &response.EventResponse{
		ID:               event.ID,
		Title:            event.Title,
		Details:          event.Details,
		Slug:             event.Slug,
		MaximumAttendees: event.MaximumAttendees,
		CreatedAt:        event.CreatedAt,
		UpdatedAt:        event.UpdatedAt,
	}, nil
}
