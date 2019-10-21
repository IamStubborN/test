package app

import (
	"github.com/IamStubborN/test/config"
	"github.com/IamStubborN/test/pkg/api/service"
	"github.com/IamStubborN/test/worker"
	"github.com/go-chi/chi"
)

func initializeAPIWorker(cfg *config.Config, router chi.Router) worker.Worker {
	return service.NewAPIWorker(
		cfg.API.Port,
		cfg.API.WriteTimeout,
		cfg.API.ReadTimeout,
		cfg.API.GracefulTimeout,
		router)
}
