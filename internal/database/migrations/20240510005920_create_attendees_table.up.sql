CREATE TABLE attendees (
   id SERIAL PRIMARY KEY, 
   name VARCHAR(250) NOT NULL,
   email VARCHAR(250) NOT NULL,
   created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
   updated_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
   event_id UUID NOT NULL REFERENCES events(id), 
   UNIQUE (email, event_id)
);