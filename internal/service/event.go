package service

import (
	"context"
	"log/slog"

	"github.com/raphaelmb/go-passin/internal/entity"
	"github.com/raphaelmb/go-passin/internal/handler/dto"
	"github.com/raphaelmb/go-passin/internal/repository"
)

type EventService interface {
	CreateEvent(ctx context.Context, e dto.EventDTO) error
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
