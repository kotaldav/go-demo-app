package main

import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

type Book struct{
  ID      string  `json:"id,omitempty"`
  Name    string  `json:"name,omitempty"`
  Author  string  `json:"author,omitempty"`
  Pages   int     `json:"pagesomitempty"`
}

var books []Book

func main() {

  books = append(books, Book{ID: "1", Name: "Kubernetes up and running", Author: "Kelsey Hightower", Pages: 192})
  books = append(books, Book{ID: "2", Name: "Site Reliability Engineering", Author: "Betsy Beyer", Pages: 552})
  books = append(books, Book{ID: "3", Name: "Docker: Up & Running", Author: "Sean P Kane", Pages: 347})

  router := mux.NewRouter()
  router.HandleFunc("/", GetRoot).Methods("GET")
  router.HandleFunc("/cpu", GetCpu).Methods("GET")
  router.HandleFunc("/kube", GetKube).Methods("GET")
  router.HandleFunc("/books", GetAllBooks).Methods("GET")
  router.HandleFunc("/books/{id}", GetBook).Methods("GET")
  router.HandleFunc("/books/{id}", CreateBook).Methods("POST")
  router.HandleFunc("/books/{id}", DeleteBook).Methods("DELETE")
  log.Fatal(http.ListenAndServe(":8000", router))
}

func GetRoot(w http.ResponseWriter, r *http.Request) {
}

func GetCpu(w http.ResponseWriter, r *http.Request) {
}

func GetKube(w http.ResponseWriter, r *http.Request) {
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
}

