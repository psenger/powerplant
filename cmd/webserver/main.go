package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"log"
)

type Book struct {
	Name string `json:"name"`
	Author string `json:"author"`
	ISBN string `json:"isbn"`
}

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func CreateBook (w http.ResponseWriter, r *http.Request)  {
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	fmt.Println(book)

}
func ReadBook (w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	isbn := vars["isbn"]
	nBook := Book {
		Name: "John",
		Author:  "Doe",
		ISBN: isbn,
	}
	json.NewEncoder(w).Encode(nBook)
}

func main() {

	//http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Welcome to my website!")
	//})
	//http.ListenAndServe(":8080", nil)

	r := mux.NewRouter()

	r.HandleFunc("/books/", logging(CreateBook)).Methods("POST")
	r.HandleFunc("/books/{isbn}", logging(ReadBook)).Methods("GET")

	http.ListenAndServe(":8080", r)
}
