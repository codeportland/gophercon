package webserver

import (
	"context"
	"net"
	"net/http"
)

type WebServer struct {
	http.Server
}

func New(host, port string, h http.Handler) *WebServer {
	var ws WebServer

	ws.Addr = net.JoinHostPort(host, port)
	ws.Handler = h

	return &ws
}

func (s *WebServer) Start() error {
	return s.ListenAndServe()
}

func (s *WebServer) Stop() error {
	return s.Shutdown(context.Background())
}
