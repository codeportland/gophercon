package main

import (
	"log"
	"os"

	"github.com/ccp-codex/gophercon/pkg/routing"
	"github.com/ccp-codex/gophercon/pkg/webserver"
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
	ws := webserver.New("", port, r)

	go func() {
		log.Fatal(ws.Start())
	}()

	internalPort := os.Getenv("INTERNAL_PORT")
	if len(internalPort) == 0 {
		log.Fatal("Internal port is not defined")
	}

	diagnosticsRouter := routing.DiagnosticsRouter()
	diagnosticsServer := webserver.New("", internalPort, diagnosticsRouter)

	log.Fatal(diagnosticsServer.Start())
}
