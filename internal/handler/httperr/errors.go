package httperr

import "errors"

var (
	ErrEventWithSameSlugFound        = errors.New("event with same slug is registered already")
	ErrEmailAlreadyRegisteredToEvent = errors.New("e-mail is already registered for this event")
	ErrMaxNumberOfAttendees          = errors.New("maximum number of attendees for this event has been reached")
	ErrEventNotFound                 = errors.New("event with given id not found")
	ErrAttendeeNotFound              = errors.New("attendee with given id not found")
	ErrAttendeeAlreadyCheckedIn      = errors.New("attendee already checked in")
)
