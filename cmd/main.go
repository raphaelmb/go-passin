package main

import (
	"log"
	"net/http"

	"github.com/raphaelmb/go-passin/internal/database"
)

func main() {
	db, err := database.NewDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	if err := http.ListenAndServe(":3333", mux); err != nil {
		log.Fatal(err)
	}
}
