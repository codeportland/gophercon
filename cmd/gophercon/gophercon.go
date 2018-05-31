package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// go run ./cmd/gophercon/gophercon.go
func main() {
	log.Printf("Service is starting...")

	r := mux.NewRouter()
	r.HandleFunc("/home", homeHandler()).Methods(http.MethodGet)
	http.ListenAndServe(":8000", r)
}

func homeHandler() func(w http.ResponseWriter, _ *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		log.Printf("Request is processing...")
		fmt.Fprint(w, "Hello! Your request was processed.")
	}
}
