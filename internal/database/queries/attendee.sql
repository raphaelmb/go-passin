-- name: GetAttendeeByEmail :one
SELECT * FROM attendees WHERE email = $1 AND event_id = $2;

-- name: CountAttendeesByEvent :one
SELECT COUNT(*) FROM attendees WHERE event_id = $1;
