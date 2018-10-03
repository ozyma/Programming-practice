package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book struct (model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//Init books variable as a slice Book struct
var books []Book

//Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	//formats the HTTP headers we will send back when we write our response
	w.Header().Set("Content-Type", "application/json")
	//use the mux.Vars() function to grab the paramters (bracketed {id} at the end of the URL we request)
	params := mux.Vars(r)
	//loop through books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

//Create a new book function
func createBook(w http.ResponseWriter, r *http.Request) {
	//formats the HTTP headers we will send back when we write our response
	w.Header().Set("Content-Type", "application/json")
	//new variable, of the custom type Book that we made (the struct above)
	var book Book
	//takes the request sent to us (in json) and Decodes it into our &book variable we declared
	_ = json.NewDecoder(r.Body).Decode(&book)
	//randomly generates an id for our book up to the number 100000000, then applies to the "ID" of our book variable (mock id - not safe)
	book.ID = strconv.Itoa(rand.Intn(100000000))
	//takes all the data we've randomly generated and Decoded from the JSON request sent, and appends it into our []books array of Book type objects
	books = append(books, book)
	//Writes a response (variable "w") to the requester with a JSON encoding of the book variable we created through this process
	json.NewEncoder(w).Encode(book)
}

// Update book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

//Delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	//formats the HTTP headers we will send back when we write our response
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)

}

func main() {
	fmt.Println("Testing out the gorilla package with Traversy media")
	//initialize our new router as variable "r"
	r := mux.NewRouter()

	//Mock Data @todo implement DB at a future point
	books = append(books, Book{ID: "1", Isbn: "606239872", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "465756845", Title: "Book two", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "3", Isbn: "542845026", Title: "Book three", Author: &Author{Firstname: "John", Lastname: "Doe"}})

	//route URL handlers AKA API Endpoints to their respective underlying functions in our Go program
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
