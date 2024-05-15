-- name: GetAttendeeByEmail :one
SELECT * FROM attendees WHERE email = $1 AND event_id = $2;

-- name: CountAttendeesByEvent :one
SELECT COUNT(*) FROM attendees WHERE event_id = $1;

-- name: GetAttendeeBadge :one
SELECT a.*, e.title FROM attendees a
JOIN events e
ON a.event_id = e.id
WHERE a.id = $1;