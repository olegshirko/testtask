// Package app configures and runs application.
package app

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"os"
	"os/signal"
	"syscall"
	"testTask/config"
	v1 "testTask/internal/controller/http/v1"
	"testTask/internal/usecase"
	"testTask/internal/usecase/repository"
	"testTask/pkg/httpserver"
	"testTask/pkg/logger"
	"testTask/pkg/postgres"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Use case
	assetUseCase := usecase.New(
		repository.New(pg),
	)

	// HTTP Server
	handler := chi.NewRouter()
	v1.NewRouter(l, assetUseCase, handler)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

}
