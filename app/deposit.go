package app

import (
	"github.com/IamStubborN/test/pkg/deposit"
	"github.com/IamStubborN/test/pkg/deposit/delivery/http"
	"github.com/IamStubborN/test/pkg/logger"
	"github.com/IamStubborN/test/pkg/responder"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
)

func registerDepositHandlers(pool *sqlx.DB, router chi.Router, l logger.Logger, r responder.Responder, duc deposit.UseCase) {
	http.RegisterDepositHandler(router, l, duc, r)
}
