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

func GetBookByName(session *gorm.DB, name string) (db.Book, error) {
	var book db.Book
	session.Where("name = ?", name).First(&book)
	if book.Name == "" {
		return db.Book{}, errors.New("error: Book not found")
	}
	return book, nil
}

func GetBooksByAuthor(session *gorm.DB, author string) ([]db.Book, error) {

	books := []db.Book{}
	session.Where("author = ?", author).Find(&books)
	if len(books) == 0 {
		return books, errors.New("error: Author not found or has no books")
	}
	return books, nil
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

func handleGetBookByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	session, err := db.GetDB()
	if err != nil {
		log.Printf("Error retrieving session : %v", err)
		http.Error(w, "Error with db", http.StatusInternalServerError)
		return
	}
	bookName := r.PathValue("name")
	book, err := GetBookByName(session, bookName)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func handleGetBooksByAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	session, err := db.GetDB()
	if err != nil {
		log.Printf("Error retrieving session : %v", err)
		http.Error(w, "Error with db", http.StatusInternalServerError)
		return
	}
	author := r.PathValue("author")
	books, err := GetBooksByAuthor(session, author)
	if err != nil {
		http.Error(w, "Author has no books or doesn't exist", http.StatusNotFound)
		return
	}
	for _, book := range books {
		json.NewEncoder(w).Encode(book)
	}
}
