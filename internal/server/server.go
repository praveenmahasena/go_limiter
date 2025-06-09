// Package server ...
package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/praveenmahasena/go-limiter/internal/algorithm"
	"github.com/praveenmahasena/go-limiter/internal/config"
)

// Server ...
type Server struct {
	listenAddr string
	*config.Config
	a   algorithm.Algorithm
	mux *http.ServeMux
}

// New ...
func New(conf *config.Config, a algorithm.Algorithm) *Server {
	return &Server{conf.GoLimiterPort, conf, a, nil}
}

// Run ...
func (s *Server) Run(ctx context.Context, errCh chan error) {
	s.mux = http.NewServeMux()
	ser := http.Server{
		Handler:      s.mux,
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
	for _, elem := range s.Rules {
		path := fmt.Sprintf("%v %v", strings.ToUpper(elem.HTTPMethod), elem.Path)
		s.mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			req, reqErr := http.NewRequest(r.Method, config.ServerAddr+r.RequestURI, nil)
			if reqErr != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{"message": "internal server error"})
				return
			}
			res, resErr := http.DefaultClient.Do(req)
			if resErr != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{"message": "internal server error"})
				return
			}
			defer res.Body.Close()
			d, err := io.ReadAll(res.Body)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{"message": "internal server error"})
				return
			}
			json.NewEncoder(w).Encode(d)
		})
	}
}
