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

type EventHandler interface {
	CreateEvent(w http.ResponseWriter, r *http.Request)
	GetEventByID(w http.ResponseWriter, r *http.Request)
}

func NewEventHandler(service service.EventService) EventHandler {
	return &handler{
		service: service,
	}
}

type handler struct {
	service service.EventService
}

func (h *handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
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
		slog.Error("error to create event", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		msg := httperr.NewInternalServerError("error decoding body")
		json.NewEncoder(w).Encode(msg)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *handler) GetEventByID(w http.ResponseWriter, r *http.Request) {
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
