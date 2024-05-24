package handler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/raphaelmb/go-passin/internal/handler/dto"
	"github.com/raphaelmb/go-passin/internal/handler/httperr"
	"github.com/raphaelmb/go-passin/internal/handler/util"
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
		util.JSONResponse(w, http.StatusBadRequest, httperr.BadRequestError("body is required"))
		return
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		slog.Error("error to decode body", "err", err)
		util.JSONResponse(w, http.StatusBadRequest, httperr.BadRequestError("error decoding body"))
		return
	}

	httpErr := httperr.ValidateHttpData(req)
	if httpErr != nil {
		slog.Error(fmt.Sprintf("error validating data: %v", httpErr))
		util.JSONResponse(w, httpErr.Code, httpErr)
		return
	}

	err = h.service.CreateEvent(r.Context(), req)
	if err != nil {
		if err == httperr.ErrEventWithSameSlugFound {
			util.JSONResponse(w, http.StatusBadRequest, httperr.BadRequestError(err.Error()))
			return
		}
		slog.Error("error to create event", "err", err)
		util.JSONResponse(w, http.StatusInternalServerError, httperr.InternalServerError("error to create event"))
		return
	}

	util.Response(w, http.StatusCreated)
}

func (h *EventHandler) GetEventByID(w http.ResponseWriter, r *http.Request) {
	stringId := r.PathValue("id")
	if stringId == "" {
		slog.Error("id not provided")
		util.JSONResponse(w, http.StatusBadRequest, httperr.BadRequestError("error to getting event, id not provided"))
		return
	}

	id, err := uuid.Parse(stringId)
	if err != nil {
		slog.Error(fmt.Sprintf("error parsing id: %v", err))
		util.JSONResponse(w, http.StatusBadRequest, httperr.BadRequestError("error parsing id"))
		return
	}

	event, err := h.service.GetEventByID(r.Context(), id)
	if err == httperr.ErrEventNotFound {
		slog.Error("error getting event with id: %v", err)
		util.JSONResponse(w, http.StatusBadRequest, httperr.NotFoundError(err.Error()))
		return
	}
	if err != nil {
		slog.Error("error getting event with id: %v", err)
		util.JSONResponse(w, http.StatusInternalServerError, httperr.InternalServerError("error getting event with id"))
		return
	}

	util.JSONResponse(w, http.StatusOK, event)
}

func (h *EventHandler) RegisterForEvent(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterForEventDTO

	if r.Body == http.NoBody {
		slog.Error("body is empty")
		util.JSONResponse(w, http.StatusBadRequest, httperr.BadRequestError("body is required"))
		return
	}

	stringId := r.PathValue("id")
	if stringId == "" {
		slog.Error("id not provided")
		util.JSONResponse(w, http.StatusBadRequest, httperr.BadRequestError("error to getting event, id not provided"))
		return
	}

	id, err := uuid.Parse(stringId)
	if err != nil {
		slog.Error(fmt.Sprintf("error parsing id: %v", err))
		util.JSONResponse(w, http.StatusBadRequest, httperr.BadRequestError("error parsing id"))
		return
	}

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		slog.Error("error to decode body", "err", err)
		util.JSONResponse(w, http.StatusBadRequest, httperr.BadRequestError("error decoding body"))
		return
	}

	httpErr := httperr.ValidateHttpData(req)
	if httpErr != nil {
		slog.Error(fmt.Sprintf("error validating data: %v", err))
		util.JSONResponse(w, httpErr.Code, httpErr)
		return
	}

	err = h.service.RegisterForEvent(r.Context(), req, id)
	if err != nil {
		switch err {
		case httperr.ErrEmailAlreadyRegisteredToEvent:
			slog.Error("error to create event", "err", err)
			util.JSONResponse(w, http.StatusBadRequest, httperr.BadRequestError(err.Error()))
			return
		case httperr.ErrMaxNumberOfAttendees:
			slog.Error("error to create event", "err", err)
			util.JSONResponse(w, http.StatusBadRequest, httperr.BadRequestError(err.Error()))
			return
		default:
			slog.Error("error to create event", "err", err)
			util.JSONResponse(w, http.StatusInternalServerError, httperr.InternalServerError("error to create event"))
			return
		}
	}

	util.Response(w, http.StatusCreated)
}
