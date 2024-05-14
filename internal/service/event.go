package service

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/gosimple/slug"

	"github.com/google/uuid"
	"github.com/raphaelmb/go-passin/internal/entity"
	"github.com/raphaelmb/go-passin/internal/handler/dto"
	"github.com/raphaelmb/go-passin/internal/handler/httperr"
	"github.com/raphaelmb/go-passin/internal/repository"
)

type EventService interface {
	CreateEvent(ctx context.Context, e dto.EventDTO) error
	GetEventByID(ctx context.Context, id uuid.UUID) (*dto.EventResponseDTO, error)
	RegisterForEvent(ctx context.Context, e dto.RegisterForEventDTO, id uuid.UUID) error
}

type EventSvc struct {
	EventRepo    repository.EventRepository
	AttendeeRepo repository.AttendeeRepository
}

func NewEventService(eventRepo repository.EventRepository, attendeeRepo repository.AttendeeRepository) EventService {
	return &EventSvc{
		EventRepo:    eventRepo,
		AttendeeRepo: attendeeRepo,
	}
}

func (s *EventSvc) CreateEvent(ctx context.Context, e dto.EventDTO) error {
	event, err := s.EventRepo.GetEventBySlug(ctx, slug.Make(e.Title))
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if event != nil {
		return httperr.ErrEventWithSameSlugFound
	}

	err = s.EventRepo.CreateEvent(ctx, &entity.Event{
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

func (s *EventSvc) GetEventByID(ctx context.Context, id uuid.UUID) (*dto.EventResponseDTO, error) {
	event, err := s.EventRepo.GetEventByID(ctx, id)
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

func (s *EventSvc) RegisterForEvent(ctx context.Context, e dto.RegisterForEventDTO, id uuid.UUID) error {
	attendee, err := s.AttendeeRepo.GetAttendeeByEmail(ctx, e.Email, id)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if attendee != nil {
		return httperr.ErrEmailAlreadyRegisteredToEvent
	}

	event, err := s.EventRepo.GetEventByID(ctx, id)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	amountOfAttendeesForEvent, err := s.AttendeeRepo.CountAttendeesByEvent(ctx, id)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if *amountOfAttendeesForEvent >= event.MaximumAttendees {
		return httperr.ErrMaxNumberOfAttendees
	}

	err = s.EventRepo.RegisterForEvent(ctx, e, id)
	if err != nil {
		return err
	}
	return nil
}
