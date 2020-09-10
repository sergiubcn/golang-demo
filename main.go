package main

import (
    "encoding/json"
    "log"
    "net/http"
    // "math/rand"
    // "strconv"
    "github.com/gorilla/mux"
)

type Book struct {
    ID string `json:"id"`
    Isbn string `json:"isbn"`
    Title string `json:"title"`
    Author *Author `json:"author"`
}

type Author struct {
    Firstname string `json:"firstname"`
    Lastname string `json:"lastname"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "application/json")
    json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "application/json")
    params := mux.Vars(r)
    for _, item := range books {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }

    json.NewEncoder(w).Encode(&Book{})
}

func main() {
    r := mux.NewRouter()

    // Mock data
    books = append(books, Book{ID: "1", Isbn: "111", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
    books = append(books, Book{ID: "2", Isbn: "222", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

    r.HandleFunc("/api/books", getBooks).Methods("GET")
    r.HandleFunc("/api/books/{id}", getBook).Methods("GET")

    log.Fatal(http.ListenAndServe(":9090", r))
    
}