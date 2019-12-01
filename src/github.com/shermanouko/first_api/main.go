package main

import(
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)


// Book struct (Model)
type Book struct{
	ID string `json: "id"`
	Isbn string `json: "isbn"`
	Title string `json: "title"`
	Author *Author `json: "author"`
}

//Author struct
type Author struct{
	FirstName string `json: "firstname`
	LastName string `json: "lastname"`
}

//init book variable as a slice book struct
var books []Book 

// get all books
func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}


// get single book
func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params
	// loop through books and find with id

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Book{})
}

//create book
func createBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(&book)

}

//update book
func updateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	parames := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)

			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(&book)
			return
		}
	}
}

//delete book
func deleteBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params
	// loop through books and find with id

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}


func main(){
	//Init router
	router := mux.NewRouter()

	//mock data
	books = append(books, Book{ID: "1", Isbn: "432535", Title: "Book One", Author: &Author{FirstName: "John", LastName: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "432536", Title: "Book Two", Author: &Author{FirstName: "Steve", LastName: "Smith"}})
	books = append(books, Book{ID: "3", Isbn: "432537", Title: "Book Three", Author: &Author{FirstName: "Mary", LastName: "Jane"}})

	//Route handlers / endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

