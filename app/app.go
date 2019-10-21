package app

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/IamStubborN/test/config"
	"github.com/IamStubborN/test/pkg/logger"
	"github.com/IamStubborN/test/worker"
)

type App struct {
	logger  logger.Logger
	Workers []worker.Worker
}

func NewApp() *App {
	var app App

	cfg := config.LoadConfig()

	app.logger = initializeLogger(cfg)
	pool := initializeSQLConn(cfg, app.logger)
	responder := initializeResponder(app.logger)
	mw := initializeMiddleWares(app.logger, responder)

	uuc := initializeUserUC(pool, app.logger)
	duc := initializeDepositUC(pool, app.logger, uuc)
	tuc := initializeTransactionUC(pool, app.logger, uuc)

	router := initializeRouter(pool, app.logger, responder, mw, uuc, duc, tuc)

	app.Workers = append(app.Workers, initializeAPIWorker(cfg, router))
	app.Workers = append(app.Workers, initializeBackupDaemon(cfg, app.logger, uuc, duc, tuc))

	return &app
}

func (app *App) Run() {
	wg := &sync.WaitGroup{}

	ctx, cancel := context.WithCancel(context.Background())

	for _, service := range app.Workers {
		wg.Add(1)
		go func(service worker.Worker) {
			defer wg.Done()
			service.Run(ctx)
		}(service)
	}

	gracefulShutdown(cancel)
	wg.Wait()
}

func gracefulShutdown(cancel context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	<-c
	close(c)
	cancel()
}
