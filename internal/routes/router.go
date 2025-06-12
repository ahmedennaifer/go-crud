package routes

import (
	"log"
	"net/http"

	mw "crud/internal/middleware"
)

func Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /book/{id}", mw.LoggingMiddleware(handleGetBookByID))
	mux.HandleFunc("GET /book/name/{name}", mw.LoggingMiddleware(handleGetBookByName))
	mux.HandleFunc("GET /book/author/{author}", mw.LoggingMiddleware(handleGetBooksByAuthor))
	mux.HandleFunc("POST /book/", mw.LoggingMiddleware(handleAddBook))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	log.Print("Server listening on port 8080...")
}
