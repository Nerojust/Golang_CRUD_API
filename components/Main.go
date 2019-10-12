package main

import (
	"NewTest/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

//init books variable as a slice Book struct
var books []models.Books

func main() {
	book1 := models.Books{
		ID:     "1",
		Title:  "Good day",
		ISBN:   "8594",
		Author: &models.Author{"Burundi", "keep"},
	}
	book2 := models.Books{
		ID:     "2",
		Title:  "Chairman dey dere day",
		ISBN:   "42494",
		Author: &models.Author{"Shenigin", "Joshua"},
	}
	book3 := models.Books{
		ID:     "3",
		Title:  "Ikase multipurpose area",
		ISBN:   "81135",
		Author: &models.Author{"Niche", "algae"},
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
	var book models.Books
	_ = json.NewDecoder(request.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // mock id
	books = append(books, book)
	json.NewEncoder(writer).Encode(book)
}

//update a book
func updateBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	parameterFromRequest := mux.Vars(request) //get the parameters
	for index, item := range books {
		if item.ID == parameterFromRequest["id"] {
			books = append(books[:index], books[index+1:]...)

			var book models.Books
			_ = json.NewDecoder(request.Body).Decode(&book)
			book.ID = parameterFromRequest["id"]
			books = append(books, book)
			json.NewEncoder(writer).Encode(book)
			return
		}
	}
	json.NewEncoder(writer).Encode(books)

}

//delete a book
func deleteBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	parameterFromRequest := mux.Vars(request) //get the parameters
	for index, item := range books {
		if item.ID == parameterFromRequest["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(writer).Encode(books)
}
