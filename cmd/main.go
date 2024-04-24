package main

import (
	"log"
	"net/http"

	"github.com/raphaelmb/go-passin/internal/database"
	"github.com/raphaelmb/go-passin/internal/database/sqlc"
	"github.com/raphaelmb/go-passin/internal/handler"
	"github.com/raphaelmb/go-passin/internal/repository"
	"github.com/raphaelmb/go-passin/internal/service"
)

func main() {
	db, err := database.NewDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	queries := sqlc.New(db)

	eventRepo := repository.NewEventRepository(queries)
	eventService := service.NewEventService(eventRepo)
	eventHandler := handler.NewEventHandler(eventService)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /events", eventHandler.CreateEvent)

	if err := http.ListenAndServe(":3333", mux); err != nil {
		log.Fatal(err)
	}
}
