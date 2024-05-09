package service

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/gosimple/slug"

	"github.com/google/uuid"
	"github.com/raphaelmb/go-passin/internal/entity"
	"github.com/raphaelmb/go-passin/internal/handler/dto"
	"github.com/raphaelmb/go-passin/internal/repository"
)

type EventService interface {
	CreateEvent(ctx context.Context, e dto.EventDTO) error
	GetEventByID(ctx context.Context, id uuid.UUID) (*dto.EventResponseDTO, error)
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
	event, err := s.repo.GetEventBySlug(ctx, slug.Make(e.Title))
	if err != nil {
		return err
	}

	if event != nil {
		return fmt.Errorf("another event with same slug already registered")
	}

	err = s.repo.CreateEvent(ctx, &entity.Event{
		Title:            e.Title,
		Details:          e.Details,
		Slug:             slug.Make(e.Title),
		MaximumAttendees: e.MaximumAttendees,
	})
	if err != nil {
		slog.Error("error creating event", "err", err)
		return err
	}

	return nil
}

func (s *service) GetEventByID(ctx context.Context, id uuid.UUID) (*dto.EventResponseDTO, error) {
	event, err := s.repo.GetEventByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &dto.EventResponseDTO{
		ID:               event.ID,
		Title:            event.Title,
		Details:          event.Details,
		Slug:             event.Slug,
		MaximumAttendees: event.MaximumAttendees,
		CreatedAt:        event.CreatedAt,
		UpdatedAt:        event.UpdatedAt,
	}, nil
}
