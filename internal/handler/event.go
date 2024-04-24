package handler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/raphaelmb/go-passin/internal/handler/dto"
	"github.com/raphaelmb/go-passin/internal/handler/httperr"
	"github.com/raphaelmb/go-passin/internal/service"
)

type EventHandler interface {
	CreateEvent(w http.ResponseWriter, r *http.Request)
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
}
