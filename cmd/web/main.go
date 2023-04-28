package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	mux := routes()

	log.Printf("Starting on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
