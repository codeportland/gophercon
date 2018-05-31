package main

import (
	"log"
	"net/http"

	"github.com/ccp-codex/gophercon/pkg/routing"
)

// go run ./cmd/gophercon/gophercon.go
func main() {
	log.Printf("Service is starting...")

	r := routing.BaseRouter()

	log.Fatal(http.ListenAndServe(":8000", r))
}
