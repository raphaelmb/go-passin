# pass.in

Pass.in is an application for **managing participants in in-person events**.

The tool allows the organizer to register an event and open a public registration page.

Registered participants can generate a credential for check-in on the day of the event.

The system will scan the participant's credential to allow entry to the event.

## Requirements

### Functional Requirements

- [ ] The organizer must be able to register a new event;
- [ ] The organizer must be able to view event data;
- [ ] The organizer must be able to view the list of participants;
- [ ] The participant must be able to register for an event;
- [ ] The participant must be able to view their registration badge;
- [ ] The participant must be able to check-in at the event;

### Business Rules

- [ ] The participant can only register for an event once;
- [ ] The participant can only register for events with available slots;
- [ ] The participant can only check-in at an event once;

### Non-functional Requirements

- [ ] Check-in at the event will be done through a QR code;