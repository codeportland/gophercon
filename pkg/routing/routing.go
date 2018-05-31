package routing

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// BaseRouter returns only business valuable routes
func BaseRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/home", homeHandler()).Methods(http.MethodGet)
	return r
}

func homeHandler() func(w http.ResponseWriter, _ *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		log.Printf("Request is processing...")
		fmt.Fprint(w, "Hello! Your request was processed.")
	}
}
