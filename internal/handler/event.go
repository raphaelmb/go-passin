package handler

import (
	"net/http"

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

func (h *handler) CreateEvent(w http.ResponseWriter, r *http.Request) {}
