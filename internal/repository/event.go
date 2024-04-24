package repository

import (
	"github.com/raphaelmb/go-passin/internal/database/sqlc"
)

type EventRepository interface {
	CreateEvent() error
}

type repository struct {
	queries *sqlc.Queries
}

func NewEventRepository(q *sqlc.Queries) EventRepository {
	return &repository{
		queries: q,
	}
}

func (r *repository) CreateEvent() error {
	return nil
}
