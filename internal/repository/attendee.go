package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/raphaelmb/go-passin/internal/database/sqlc"
	"github.com/raphaelmb/go-passin/internal/entity"
)

type AttendeeRepository interface {
	GetAttendeeByEmail(ctx context.Context, email string, eventID uuid.UUID) (*entity.Attendee, error)
	CountAttendeesByEvent(ctx context.Context, id uuid.UUID) (*int, error)
	GetAttendeeBadge(ctx context.Context, id int) (*entity.Attendee, error)
	GetCheckIn(ctx context.Context, id int) (*entity.CheckIn, error)
	CreateCheckIn(ctx context.Context, id int) error
}

type AttendeeRepo struct {
	queries *sqlc.Queries
}

func NewAttendeeRepository(q *sqlc.Queries) AttendeeRepository {
	return &AttendeeRepo{
		queries: q,
	}
}

func (r *AttendeeRepo) GetAttendeeByEmail(ctx context.Context, email string, eventID uuid.UUID) (*entity.Attendee, error) {
	attendee, err := r.queries.GetAttendeeByEmail(ctx, sqlc.GetAttendeeByEmailParams{Email: email, EventID: eventID})
	if err != nil {
		return nil, err
	}
	return &entity.Attendee{
		ID:    int(attendee.ID),
		Email: attendee.Email,
		Name:  attendee.Name,
	}, nil
}

func (r *AttendeeRepo) CountAttendeesByEvent(ctx context.Context, id uuid.UUID) (*int, error) {
	num, err := r.queries.CountAttendeesByEvent(ctx, id)
	if err != nil {
		return nil, err
	}

	result := int(num)
	return &result, nil
}

func (r *AttendeeRepo) GetAttendeeBadge(ctx context.Context, id int) (*entity.Attendee, error) {
	attendee, err := r.queries.GetAttendeeBadge(ctx, int32(id))
	if err != nil {
		return nil, err
	}
	return &entity.Attendee{
		ID:         int(attendee.ID),
		Name:       attendee.Name,
		Email:      attendee.Email,
		EventTitle: attendee.Title,
	}, nil
}

func (r *AttendeeRepo) GetCheckIn(ctx context.Context, id int) (*entity.CheckIn, error) {
	attendeeCheckIn, err := r.queries.GetAttendeeByID(ctx, int32(id))
	if err != nil {
		return nil, err
	}
	return &entity.CheckIn{
		ID:         int(attendeeCheckIn.ID),
		CreatedAt:  attendeeCheckIn.CreatedAt,
		AttendeeID: int(attendeeCheckIn.AttendeeID),
	}, nil
}

func (r *AttendeeRepo) CreateCheckIn(ctx context.Context, id int) error {
	err := r.queries.CreateCheckIn(ctx, int32(id))
	if err != nil {
		return err
	}
	return nil
}
