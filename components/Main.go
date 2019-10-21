package main

import (
	"NewTest/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
)

//init totalBooksArraySlice variable as a slice Book struct
var totalBooksArraySlice []models.Books

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
	totalBooksArraySlice = append(totalBooksArraySlice, book1, book2, book3)

	router := mux.NewRouter()
	//route handler for endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}

//get all totalBooksArraySlice
func getBooks(writer http.ResponseWriter, request *http.Request) {
	//model a new response plus the data slice
	var bookRes = models.Response{
		Message: models.SUCCESS,
		Success: true,
		Data:    totalBooksArraySlice,
	}
	writer.Header().Set("Content-Type", "application/json")
	//encode and return the full object response
	json.NewEncoder(writer).Encode(bookRes)
}

//get a single book
func getBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	parameterFromRequest := mux.Vars(request) //get the parameters
	//loop through totalBooksArraySlice and find the id
	for _, singleBookItem := range totalBooksArraySlice {
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
	// a single newBookRequestFromClient model reference
	var newBookRequestFromClient models.Books
	//decode the incoming body and decode the new newBookRequestFromClient
	_ = json.NewDecoder(request.Body).Decode(&newBookRequestFromClient)
	//generate a new id for the new newBookRequestFromClient
	newBookRequestFromClient.ID = strconv.Itoa(rand.Intn(10000000)) // mock id
	//add the new one to the existing list
	totalBooksArraySlice = append(totalBooksArraySlice, newBookRequestFromClient)
	//encode and show the
	json.NewEncoder(writer).Encode("Added Successfully")
	//json.NewEncoder(writer).Encode(newBookRequestFromClient)
}

//update a book
func updateBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	parameterFromRequest := mux.Vars(request) //get the parameters
	for index, item := range totalBooksArraySlice {
		if item.ID == parameterFromRequest["id"] {
			totalBooksArraySlice = append(totalBooksArraySlice[:index], totalBooksArraySlice[index+1:]...)
			var book models.Books
			_ = json.NewDecoder(request.Body).Decode(&book)
			book.ID = parameterFromRequest["id"]
			totalBooksArraySlice = append(totalBooksArraySlice, book)
			json.NewEncoder(writer).Encode(book)
			return
		}
	}
	json.NewEncoder(writer).Encode("Updated Successfully")
	//json.NewEncoder(writer).Encode(totalBooksArraySlice)
}

//delete a book
func deleteBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	parameterFromRequest := mux.Vars(request) //get the parameters
	for index, singleItem := range totalBooksArraySlice {
		if singleItem.ID == parameterFromRequest["id"] {
			totalBooksArraySlice = append(totalBooksArraySlice[:index], totalBooksArraySlice[index+1:]...)
			break
		}
	}
	//json.NewEncoder(writer).Encode(totalBooksArraySlice)
	json.NewEncoder(writer).Encode("Deleted Successfully")
}
