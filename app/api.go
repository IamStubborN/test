package app

import (
	"github.com/go-chi/chi"

	"github.com/IamStubborN/test/config"
	"github.com/IamStubborN/test/daemon"
	"github.com/IamStubborN/test/pkg/api/service"
)

func initializeAPIDaemon(cfg *config.Config, router chi.Router) daemon.Daemon {
	return service.NewAPIWorker(
		cfg.API.Port,
		cfg.API.WriteTimeout,
		cfg.API.ReadTimeout,
		cfg.API.GracefulTimeout,
		router)
}
