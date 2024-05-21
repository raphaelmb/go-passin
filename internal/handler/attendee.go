package handler

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/raphaelmb/go-passin/internal/handler/httperr"
	"github.com/raphaelmb/go-passin/internal/service"
)

type IAttendeeHandler interface {
	GetAttendeeBadge(w http.ResponseWriter, r *http.Request)
	CheckIn(w http.ResponseWriter, r *http.Request)
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

	ctx := context.WithValue(r.Context(), "req", r)

	attendee, err := h.service.GetAttendeeBadge(ctx, id)
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

func (h *AttendeeHandler) CheckIn(w http.ResponseWriter, r *http.Request) {
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

	err = h.service.CreateCheckIn(r.Context(), id)
	if err == httperr.ErrAttendeeAlreadyCheckedIn {
		slog.Error("attendee already checked in")
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError(err.Error())
		json.NewEncoder(w).Encode(msg)
		return
	}
	if err != nil {
		slog.Error("error checkin in the attendee")
		w.WriteHeader(http.StatusInternalServerError)
		msg := httperr.NewInternalServerError("error checkin in the attendee")
		json.NewEncoder(w).Encode(msg)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
