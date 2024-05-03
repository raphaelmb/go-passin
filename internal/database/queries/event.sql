-- name: CreateEvent :exec
INSERT INTO events(title, details, slug, maximum_attendees) VALUES($1, $2, $3, $4);

-- name: GetEventByID :one
SELECT * FROM events WHERE id = $1;