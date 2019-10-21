package app

import (
	"github.com/IamStubborN/test/pkg/logger"
	"github.com/IamStubborN/test/pkg/responder"
	"github.com/IamStubborN/test/pkg/transaction"
	"github.com/IamStubborN/test/pkg/transaction/delivery/http"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
)

func registerTransactionHandlers(pool *sqlx.DB, router chi.Router, l logger.Logger, r responder.Responder, tuc transaction.UseCase) {
	http.RegisterTransactionHandler(router, l, tuc, r)
}

