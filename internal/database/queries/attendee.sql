-- name: GetAttendeeByEmail :one
SELECT * FROM attendees WHERE email = $1 AND event_id = $2;

-- name: CountAttendeesByEvent :one
SELECT COUNT(*) FROM attendees WHERE event_id = $1;

-- name: GetAttendeeBadge :one
SELECT a.*, e.title FROM attendees a
JOIN events e
ON a.event_id = e.id
WHERE a.id = $1;

-- name: GetAttendeeByID :one
SELECT * From checkins WHERE attendee_id = $1;

-- name: CreateCheckIn :exec
INSERT INTO checkins(attendee_id) VALUES($1);