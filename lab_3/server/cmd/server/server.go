package main

import (
	"context"
	"fmt"
	"net/http"
)

type HttpPortNumber int

// ChatApiServer configures necessary handlers and starts listening on a configured port.
type ApiServer struct {
	Port HttpPortNumber

	router *Router

	server *http.Server
}

// Start will set all handlers and start listening.
// If this methods succeeds, it does not return until server is shut down.
// Returned error will never be nil.
func (s *ApiServer) Start() error {

	handler := new(http.ServeMux)
	handler.HandleFunc("/", s.router.handle)

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: handler,
	}

	return s.server.ListenAndServe()
}

// Stops will shut down previously started HTTP server.
func (s *ApiServer) Stop() error {
	if s.server == nil {
		return fmt.Errorf("server was not started")
	}
	return s.server.Shutdown(context.Background())
}
