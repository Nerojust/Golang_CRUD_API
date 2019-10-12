package main

import (
	"NewTest/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//init books variable as a slice Book struct
var books []models.Books

func main() {

	book1 := models.Books{
		ID:     "1",
		Title:  "Swan princess",
		ISBN:   "2389283",
		Author: models.Author{"Gregory", "Franklin"},
	}
	book2 := models.Books{
		ID:     "2",
		Title:  "Love Boulevard",
		ISBN:   "195048",
		Author: models.Author{"Fred", "Hammond"},
	}
	book3 := models.Books{
		ID:     "3",
		Title:  "Christmas Miracle",
		ISBN:   "774893",
		Author: models.Author{"Nerojust", "Greta"},
	}
	//attach the slices
	books = append(books, book1, book2, book3)

	router := mux.NewRouter()
	//route handler for endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

//get all books
func getBooks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(books)
}

//get a single book
func getBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	parameterFromRequest := mux.Vars(request) //get the parameters
	//loop through books and find the id
	for _, singleBookItem := range books {
		if singleBookItem.ID == parameterFromRequest["id"] {
			//when u get it encode it to json
			json.NewEncoder(writer).Encode(singleBookItem)
			return
		}
	}
	//finally encode the whole struct and return it to client
	json.NewEncoder(writer).Encode(&models.Books{})
}

//create a new book
func createBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
}

//update a book
func updateBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
}

//delete a book
func deleteBook(writer http.ResponseWriter, request *http.Request) {

}
