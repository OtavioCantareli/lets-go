package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Ayooo"))
}

func view(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Viewing post %d", id)
}

func create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method now allowed.", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Creating..."))
}

func routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/view", view)
	mux.HandleFunc("/create", create)

	return mux
}

func main() {
	mux := routes()

	log.Print("Starting on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
