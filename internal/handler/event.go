package handler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/raphaelmb/go-passin/internal/handler/dto"
	"github.com/raphaelmb/go-passin/internal/handler/httperr"
	"github.com/raphaelmb/go-passin/internal/service"
)

type IEventHandler interface {
	CreateEvent(w http.ResponseWriter, r *http.Request)
	GetEventByID(w http.ResponseWriter, r *http.Request)
	RegisterForEvent(w http.ResponseWriter, r *http.Request)
}

type EventHandler struct {
	service service.EventService
}

func NewEventHandler(service service.EventService) IEventHandler {
	return &EventHandler{
		service: service,
	}
}

func (h *EventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var req dto.EventDTO

	if r.Body == http.NoBody {
		slog.Error("body is empty")
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("body is required")
		json.NewEncoder(w).Encode(msg)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		slog.Error("error to decode body", "err", err)
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error decoding body")
		json.NewEncoder(w).Encode(msg)
		return
	}

	httpErr := httperr.ValidateHttpData(req)
	if httpErr != nil {
		slog.Error(fmt.Sprintf("error validating data: %v", err))
		w.WriteHeader(httpErr.Code)
		json.NewEncoder(w).Encode(httpErr)
		return
	}

	err = h.service.CreateEvent(r.Context(), req)
	if err != nil {
		if err == httperr.ErrEventWithSameSlugFound {
			w.WriteHeader(http.StatusBadRequest)
			msg := httperr.NewBadRequestError(err.Error())
			json.NewEncoder(w).Encode(msg)
			return
		}
		slog.Error("error to create event", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		msg := httperr.NewInternalServerError("error to create event")
		json.NewEncoder(w).Encode(msg)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *EventHandler) GetEventByID(w http.ResponseWriter, r *http.Request) {
	stringId := r.PathValue("id")
	if stringId == "" {
		slog.Error("id not provided")
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error to getting event, id not provided")
		json.NewEncoder(w).Encode(msg)
		return
	}

	id, err := uuid.Parse(stringId)
	if err != nil {
		slog.Error(fmt.Sprintf("error parsing id: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error parsing id")
		json.NewEncoder(w).Encode(msg)
		return
	}

	event, err := h.service.GetEventByID(r.Context(), id)
	if err == httperr.ErrEventNotFound {
		slog.Error("error getting event with id: %v", err)
		w.WriteHeader(http.StatusNotFound)
		msg := httperr.NewNotFoundError(err.Error())
		json.NewEncoder(w).Encode(msg)
		return
	}
	if err != nil {
		slog.Error("error getting event with id: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		msg := httperr.NewInternalServerError("error getting event with id")
		json.NewEncoder(w).Encode(msg)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(event)
}

func (h *EventHandler) RegisterForEvent(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterForEventDTO

	if r.Body == http.NoBody {
		slog.Error("body is empty")
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("body is required")
		json.NewEncoder(w).Encode(msg)
		return
	}

	stringId := r.PathValue("id")
	if stringId == "" {
		slog.Error("id not provided")
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error to getting event, id not provided")
		json.NewEncoder(w).Encode(msg)
		return
	}

	id, err := uuid.Parse(stringId)
	if err != nil {
		slog.Error(fmt.Sprintf("error parsing id: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error parsing id")
		json.NewEncoder(w).Encode(msg)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		slog.Error("error to decode body", "err", err)
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error decoding body")
		json.NewEncoder(w).Encode(msg)
		return
	}

	httpErr := httperr.ValidateHttpData(req)
	if httpErr != nil {
		slog.Error(fmt.Sprintf("error validating data: %v", err))
		w.WriteHeader(httpErr.Code)
		json.NewEncoder(w).Encode(httpErr)
		return
	}

	err = h.service.RegisterForEvent(r.Context(), req, id)
	if err == httperr.ErrEmailAlreadyRegisteredToEvent {
		slog.Error("error to create event", "err", err)
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError(err.Error())
		json.NewEncoder(w).Encode(msg)
		return
	}

	if err == httperr.ErrMaxNumberOfAttendees {
		slog.Error("error to create event", "err", err)
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError(err.Error())
		json.NewEncoder(w).Encode(msg)
		return
	}

	if err != nil {
		slog.Error("error to create event", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		msg := httperr.NewInternalServerError("error to create event")
		json.NewEncoder(w).Encode(msg)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
