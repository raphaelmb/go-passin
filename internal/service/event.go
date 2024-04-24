package service

import "github.com/raphaelmb/go-passin/internal/repository"

type EventService interface {
	CreateEvent() error
}

type service struct {
	repo repository.EventRepository
}

func NewEventService(repo repository.EventRepository) EventService {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateEvent() error {
	return nil
}
