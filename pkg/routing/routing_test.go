package routing_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ccp-codex/gophercon/pkg/routing"
)

func TestBaseRouter(t *testing.T) {
	handler := routing.BaseRouter()

	ts := httptest.NewServer(handler)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/home")
	if err != nil {
		log.Fatal(err)
	}

	// todo: you can use testify library instead
	if res.StatusCode != http.StatusOK {
		t.Errorf("Wrong status code %d (%d expected)", res.StatusCode, http.StatusOK)
	}

}
