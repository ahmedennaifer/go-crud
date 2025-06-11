package routes

import (
	"log"
	"net/http"

	mw "crud/internal/middleware"
)

func Start() {
	http.HandleFunc("GET /book/{id}", mw.LoggingMiddleware(handleGetBookByID))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	log.Print("Server listening on port 8080...")
}
