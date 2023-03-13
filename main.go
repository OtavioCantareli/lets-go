package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ayooo"))
}

func view(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Viewing..."))
}

func create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/view", view)
	mux.HandleFunc("/create", create)

	log.Print("Starting on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
