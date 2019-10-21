package app

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/IamStubborN/test/config"
	"github.com/IamStubborN/test/daemon"
	"github.com/IamStubborN/test/pkg/logger"
)

type App struct {
	logger  logger.Logger
	Daemons []daemon.Daemon
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

	router := initializeRouter(app.logger, responder, mw, uuc, duc, tuc)

	app.Daemons = append(app.Daemons, initializeAPIDaemon(cfg, router))
	app.Daemons = append(app.Daemons, initializeBackupDaemon(cfg, app.logger, uuc, duc, tuc))

	return &app
}

func (app *App) Run() {
	wg := &sync.WaitGroup{}

	ctx, cancel := context.WithCancel(context.Background())

	for _, service := range app.Daemons {
		wg.Add(1)
		go func(service daemon.Daemon) {
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
