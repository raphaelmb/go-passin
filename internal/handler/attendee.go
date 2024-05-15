package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/raphaelmb/go-passin/internal/handler/httperr"
	"github.com/raphaelmb/go-passin/internal/service"
)

type IAttendeeHandler interface {
	GetAttendeeBadge(w http.ResponseWriter, r *http.Request)
}

func NewAttendeeHandler(service service.AttendeeService) IAttendeeHandler {
	return &AttendeeHandler{
		service: service,
	}
}

type AttendeeHandler struct {
	service service.AttendeeService
}

func (h *AttendeeHandler) GetAttendeeBadge(w http.ResponseWriter, r *http.Request) {
	stringId := r.PathValue("id")
	if stringId == "" {
		slog.Error("id not provided")
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error to getting event, id not provided")
		json.NewEncoder(w).Encode(msg)
		return
	}

	id, err := strconv.Atoi(stringId)
	if err != nil {
		slog.Error("error converting id string to int")
		w.WriteHeader(http.StatusInternalServerError)
		msg := httperr.NewInternalServerError("error converting id string to int")
		json.NewEncoder(w).Encode(msg)
		return
	}

	attendee, err := h.service.GetAttendeeBadge(r.Context(), id)
	if err == httperr.ErrAttendeeNotFound {
		slog.Error("attendee with given id not found")
		w.WriteHeader(http.StatusInternalServerError)
		msg := httperr.NewBadRequestError(err.Error())
		json.NewEncoder(w).Encode(msg)
		return
	}
	if err != nil {
		slog.Error("error converting id string to int")
		w.WriteHeader(http.StatusInternalServerError)
		msg := httperr.NewInternalServerError("error converting id string to int")
		json.NewEncoder(w).Encode(msg)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(attendee)
}
