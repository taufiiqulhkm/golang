package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get single book by ID
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()
	id, _ := strconv.Atoi(params.Get("id"))

	for _, book := range books {
		if book.ID == id {
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	json.NewEncoder(w).Encode(&Book{})
}

// Create a new book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	book.ID = len(books) + 1
	books = append(books, book)

	json.NewEncoder(w).Encode(book)
}

// Update an existing book by ID
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()
	id, _ := strconv.Atoi(params.Get("id"))

	for index, book := range books {
		if book.ID == id {
			var updatedBook Book
			_ = json.NewDecoder(r.Body).Decode(&updatedBook)
			updatedBook.ID = id
			books[index] = updatedBook
			json.NewEncoder(w).Encode(updatedBook)
			return
		}
	}

	json.NewEncoder(w).Encode(&Book{})
}

// Delete an existing book by ID
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()
	id, _ := strconv.Atoi(params.Get("id"))

	for index, book := range books {
		if book.ID == id {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(books)
}

func main() {
	// Sample books
	books = append(books, Book{ID: 1, Title: "Dilan 1991", Author: "Pidi Baiq"})
	books = append(books, Book{ID: 2, Title: "Dillan 1990", Author: "Pidi Baiq"})

	// Define routes
	http.HandleFunc("/books", getBooks)
	http.HandleFunc("/books/get", getBook)
	http.HandleFunc("/books/create", createBook)
	http.HandleFunc("/books/update", updateBook)
	http.HandleFunc("/books/delete", deleteBook)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
