package main

import (
	"context"
	"fmt"
	"net/http"
)

// APIServer configures necessary handlers and starts listening on a configured port.
type APIServer struct {
	Port int

	TreeHandler http.HandlerFunc

	server *http.Server
}

// Start will set all handlers and start listening.
// If this methods succeeds, it does not return until server is shut down.
// Returned error will never be nil.
func (s *APIServer) Start() error {
	if s.TreeHandler == nil {
		return fmt.Errorf("HTTP handler is not defined - cannot start")
	}
	if s.Port == 0 {
		return fmt.Errorf("port is not defined")
	}

	handler := new(http.ServeMux)
	handler.HandleFunc("/tree", s.TreeHandler)

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: handler,
	}

	return s.server.ListenAndServe()
}

// Stop will shut down previously started HTTP server.
func (s *APIServer) Stop() error {
	if s.server == nil {
		return fmt.Errorf("server was not started")
	}
	return s.server.Shutdown(context.Background())
}
