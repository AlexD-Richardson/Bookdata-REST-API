package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct{}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Get got gotten"}`))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "Post got posted"}`))
	case "PUT":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "Put got put in it's place"}`))
	case "DELETE":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "Delete didn't get deleted"}`))
	default:
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "Method not found. What are you doing?"}`))
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8080", r))
}
