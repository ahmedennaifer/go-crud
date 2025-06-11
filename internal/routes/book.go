package routes

import (
	"crud/internal/handlers/db"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

func GetBookByID(session *gorm.DB, id int) (db.Book, error) {

	var book db.Book
	session.Find(&book, id)
	if book.Name == "" {
		return db.Book{}, errors.New("error: Book not found")
	}
	return book, nil
}

func handleGetBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	session, err := db.GetDB()
	if err != nil {
		log.Printf("Error retrieving session : %v", err)
		http.Error(w, "Error with db", http.StatusInternalServerError)
		return
	}
	bookID, convErr := strconv.Atoi(r.PathValue("id"))
	if convErr != nil {
		log.Printf("Cannot convert id to int: %v", convErr)
		http.Error(w, "Malformed id", http.StatusBadRequest)
		return
	}
	book, err := GetBookByID(session, bookID)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(book)
}
