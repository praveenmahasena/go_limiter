// Package server ...
package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/praveenmahasena/go-limiter/internal/algorithm"
	"github.com/praveenmahasena/go-limiter/internal/config"
)

// Server ...
type Server struct {
	listenAddr string
	*config.Config
	algorithm.Algo
}

// New ...
func New(conf *config.Config, a algorithm.Algo) *Server {
	return &Server{conf.GoLimiterPort, conf, a}
}

// Run ...
func (s *Server) Run(ctx context.Context, errCh chan error) {
	mux := http.NewServeMux()

	ser := http.Server{
		Handler:      mux,
		Addr:         s.listenAddr,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
		IdleTimeout:  time.Second,
	}

	s.route()

	go func() {
		if err := ser.ListenAndServe(); err != nil {
			errCh <- fmt.Errorf("error during starting up server or Shutdown happened with value %v", err)
		}
	}()

	<-ctx.Done()
	if err := ser.Shutdown(context.Background()); err != nil {
		errCh <- fmt.Errorf("error during shutting down server with value %v", err)
	}
}

func (s *Server) route() {
	for _, elem := range s.Config.Rules {
		fmt.Println(elem)
	}
}
