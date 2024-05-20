package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"

	"github.com/raphaelmb/go-passin/internal/database"
	"github.com/raphaelmb/go-passin/internal/database/sqlc"
	"github.com/raphaelmb/go-passin/internal/handler"
	"github.com/raphaelmb/go-passin/internal/repository"
	"github.com/raphaelmb/go-passin/internal/service"
)

func main() {
	db, err := database.NewDBConnection()
	if err != nil {
		slog.Error("error connecting to database", "err", err)
	}
	defer db.Close()

	queries := sqlc.New(db)

	// attendees
	attendeeRepo := repository.NewAttendeeRepository(queries)
	attendeeSvc := service.NewAttendeeSvc(attendeeRepo)
	attendeeHandler := handler.NewAttendeeHandler(attendeeSvc)

	// events
	eventRepo := repository.NewEventRepository(queries)
	eventService := service.NewEventService(eventRepo, attendeeSvc)
	eventHandler := handler.NewEventHandler(eventService)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handler.GetHealth)
	mux.HandleFunc("POST /events", eventHandler.CreateEvent)
	mux.HandleFunc("GET /events/{id}", eventHandler.GetEventByID)
	mux.HandleFunc("POST /events/{id}/attendees", eventHandler.RegisterForEvent)
	mux.HandleFunc("GET /attendees/{id}/badge", attendeeHandler.GetAttendeeBadge)

	port := os.Getenv("PORT")
	slog.Info(fmt.Sprintf("server running on port %s", port))
	if err := http.ListenAndServe(port, mux); err != nil {
		slog.Error("error starting server", "err", err)
	}
}
