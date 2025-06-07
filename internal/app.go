// Package internal ...
package internal

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/praveenmahasena/go-limiter/internal/algorithm"
	"github.com/praveenmahasena/go-limiter/internal/config"
	"github.com/praveenmahasena/go-limiter/internal/server"
)

// Run ...
func Run() error {
	config, configErr := config.Load()
	if configErr != nil {
		return configErr
	}
	algo, algoErr := algorithm.New(config.Rules)
	if algoErr != nil {
		return algoErr
	}
	server := server.New(config, algo)
	ctx, cancel := context.WithCancel(context.Background())
	errCh := make(chan error)
	defer cancel()

	go func(cancel context.CancelFunc) {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		cancel()
	}(cancel)

	go func(ctx context.Context, errCh chan error) {
		server.Run(ctx, errCh)
	}(ctx, errCh)

	select {
	case <-ctx.Done():
		fmt.Println("cancel signal has been sent to server ...")
	case err := <-errCh:
		return err
	}

	return <-errCh
}
