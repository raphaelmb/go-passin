package service

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/raphaelmb/go-passin/internal/handler/dto"
	"github.com/raphaelmb/go-passin/internal/handler/httperr"
	"github.com/raphaelmb/go-passin/internal/repository"
)

type AttendeeService interface {
	GetAttendeeByEmail(ctx context.Context, email string, eventID uuid.UUID) (*dto.AttendeeResponseDTO, error)
	CountAttendeesByEvent(ctx context.Context, id uuid.UUID) (*int, error)
	GetAttendeeBadge(ctx context.Context, id int) (*dto.AttendeeResponseDTO, error)
}

type AttendeeSvc struct {
	repo repository.AttendeeRepository
}

func NewAttendeeSvc(repo repository.AttendeeRepository) AttendeeService {
	return &AttendeeSvc{
		repo: repo,
	}
}

func (s *AttendeeSvc) GetAttendeeByEmail(ctx context.Context, email string, eventID uuid.UUID) (*dto.AttendeeResponseDTO, error) {
	attendee, err := s.repo.GetAttendeeByEmail(ctx, email, eventID)
	if err != nil {
		return nil, err
	}
	return &dto.AttendeeResponseDTO{
		ID:    attendee.ID,
		Name:  attendee.Name,
		Email: attendee.Email,
	}, nil
}

func (s *AttendeeSvc) CountAttendeesByEvent(ctx context.Context, id uuid.UUID) (*int, error) {
	amount, err := s.repo.CountAttendeesByEvent(ctx, id)
	if err != nil {
		return nil, err
	}
	return amount, nil
}

func (s *AttendeeSvc) GetAttendeeBadge(ctx context.Context, id int) (*dto.AttendeeResponseDTO, error) {
	attendee, err := s.repo.GetAttendeeBadge(ctx, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, httperr.ErrAttendeeNotFound
	}

	return &dto.AttendeeResponseDTO{
		Name:       attendee.Name,
		Email:      attendee.Email,
		EventTitle: attendee.EventTitle,
	}, nil
}
