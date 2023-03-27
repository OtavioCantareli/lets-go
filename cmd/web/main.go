package main

import (
	"log"
	"net/http"
)

func main() {
	mux := routes()

	log.Print("Starting on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
