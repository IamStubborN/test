package app

import (
	"github.com/IamStubborN/test/pkg/deposit"
	"github.com/IamStubborN/test/pkg/logger"
	"github.com/IamStubborN/test/pkg/mware"
	"github.com/IamStubborN/test/pkg/responder"
	"github.com/IamStubborN/test/pkg/transaction"
	"github.com/IamStubborN/test/pkg/user"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
)

func initializeRouter(
	pool *sqlx.DB,
	l logger.Logger,
	r responder.Responder,
	mw mware.MWare,
	uuc user.UseCase,
	duc deposit.UseCase,
	tuc transaction.UseCase) chi.Router {

	router := chi.NewRouter()
	router.Use(mw.AuthMiddleware)
	router.Use(mw.RequestLogger)

	registerHandlers(pool, router, l, r, uuc, duc, tuc)

	return router
}
