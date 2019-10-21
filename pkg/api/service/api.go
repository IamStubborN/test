package service

import (
	"context"
	"github.com/IamStubborN/test/worker"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type APIWorker struct {
	Port     int
	Router   chi.Router
	WTimeout time.Duration
	RTimeout time.Duration
	GTimeout time.Duration
}

func NewAPIWorker(port int, writeTimeout, readTimeout, gracefulTimeout time.Duration, router chi.Router) worker.Worker {
	return &APIWorker{
		Port:     port,
		Router:   router,
		WTimeout: writeTimeout,
		RTimeout: readTimeout,
		GTimeout: gracefulTimeout,
	}
}

func (aw *APIWorker) Run(ctx context.Context) {
	srv := &http.Server{
		Addr:         ":" + strconv.Itoa(aw.Port),
		Handler:      chi.ServerBaseContext(ctx, aw.Router),
		WriteTimeout: aw.WTimeout,
		ReadTimeout:  aw.RTimeout,
	}

	go func() {
		<-ctx.Done()

		ctxShutDown, cancel := context.WithTimeout(context.Background(), aw.GTimeout)
		defer cancel()

		if err := srv.Shutdown(ctxShutDown); err != nil {
			zap.L().Info(err.Error())
		}
	}()

	if err := srv.ListenAndServe(); err != nil {
		zap.L().Info(err.Error())
	}
}
