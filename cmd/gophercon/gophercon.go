package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ccp-codex/gophercon/pkg/routing"
)

// go run ./cmd/gophercon/gophercon.go
func main() {
	log.Printf("Service is starting...")

	// todo: you can also use github.com/kelsey/hightower/envconfig
	// to keep your config more structured
	port := os.Getenv("SERVICE_PORT")
	if len(port) == 0 {
		log.Fatal("Service port was not defined")
	}

	r := routing.BaseRouter()

	log.Fatal(http.ListenAndServe(":"+port, r))
}
