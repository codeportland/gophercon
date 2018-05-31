package routing

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func DiagnosticsRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/healthy", healthHandler()).Methods(http.MethodGet)
	r.HandleFunc("/ready", readyHandler()).Methods(http.MethodGet)

	return r
}

func readyHandler() func(w http.ResponseWriter, _ *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, http.StatusText(http.StatusOK))
	}
}

func healthHandler() func(w http.ResponseWriter, _ *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, http.StatusText(http.StatusOK))
	}
}
