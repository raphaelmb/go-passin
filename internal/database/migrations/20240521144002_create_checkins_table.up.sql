CREATE TABLE checkins (
   id SERIAL PRIMARY KEY, 
   created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
   attendee_id INTEGER UNIQUE NOT NULL REFERENCES attendees(id)
);