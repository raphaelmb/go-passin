package httperr

import "errors"

var (
	ErrEventWithSameSlugFound        = errors.New("event with same slug is registered already")
	ErrEmailAlreadyRegisteredToEvent = errors.New("e-mail is already registered for this event")
	ErrMaxNumberOfAttendees          = errors.New("maximum number of attendees for this event has been reached")
)
