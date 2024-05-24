package handler

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/raphaelmb/go-passin/internal/handler/httperr"
	"github.com/raphaelmb/go-passin/internal/handler/util"
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
		util.JSONResponse(w, http.StatusBadRequest, httperr.BadRequestError("error to getting event, id not provided"))
		return
	}

	id, err := strconv.Atoi(stringId)
	if err != nil {
		slog.Error("error converting id string to int")
		util.JSONResponse(w, http.StatusInternalServerError, httperr.InternalServerError("error converting id string to int"))
		return
	}

	ctx := context.WithValue(r.Context(), "req", r)

	attendee, err := h.service.GetAttendeeBadge(ctx, id)
	if err == httperr.ErrAttendeeNotFound {
		slog.Error("attendee with given id not found")
		util.JSONResponse(w, http.StatusInternalServerError, httperr.BadRequestError(err.Error()))
		return
	}
	if err != nil {
		slog.Error("error converting id string to int")
		util.JSONResponse(w, http.StatusInternalServerError, httperr.InternalServerError("error converting id string to int"))
		return
	}

	util.JSONResponse(w, http.StatusOK, attendee)
}

func (h *AttendeeHandler) CheckIn(w http.ResponseWriter, r *http.Request) {
	stringId := r.PathValue("id")
	if stringId == "" {
		slog.Error("id not provided")
		util.JSONResponse(w, http.StatusBadRequest, httperr.BadRequestError("error to getting event, id not provided"))
		return
	}

	id, err := strconv.Atoi(stringId)
	if err != nil {
		slog.Error("error converting id string to int")
		util.JSONResponse(w, http.StatusInternalServerError, httperr.InternalServerError("error converting id string to int"))
		return
	}

	err = h.service.CreateCheckIn(r.Context(), id)
	if err == httperr.ErrAttendeeAlreadyCheckedIn {
		slog.Error("attendee already checked in")
		util.JSONResponse(w, http.StatusBadRequest, httperr.BadRequestError(err.Error()))
		return
	}
	if err != nil {
		slog.Error("error checkin in the attendee")
		util.JSONResponse(w, http.StatusInternalServerError, httperr.InternalServerError("error checkin in the attendee"))
		return
	}

	util.Response(w, http.StatusCreated)
}
