-- name: CreateEvent :exec
INSERT INTO events(title, details, slug, maximum_attendees) VALUES($1, $2, $3, $4);

-- name: GetEventByID :one
SELECT e.*, COUNT(a) as attendees FROM events e
JOIN attendees a
ON e.id = a.event_id
WHERE e.id = $1
GROUP BY e.id;

-- name: GetEventBySlug :one
SELECT * FROM events WHERE slug = $1;

-- name: RegisterForEvent :exec
INSERT INTO attendees(name, email, event_id) VALUES($1, $2, $3);