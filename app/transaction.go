package app

import (
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"

	"github.com/IamStubborN/test/pkg/logger"
	"github.com/IamStubborN/test/pkg/responder"
	"github.com/IamStubborN/test/pkg/transaction"
	"github.com/IamStubborN/test/pkg/transaction/cache"
	"github.com/IamStubborN/test/pkg/transaction/delivery/http"
	"github.com/IamStubborN/test/pkg/transaction/repository"
	"github.com/IamStubborN/test/pkg/transaction/usecase"
	"github.com/IamStubborN/test/pkg/user"
)

func registerTransactionHandlers(router chi.Router, l logger.Logger, r responder.Responder, tuc transaction.UseCase) {
	http.RegisterTransactionHandler(router, l, tuc, r)
}

func initializeTransactionUC(pool *sqlx.DB, l logger.Logger, uuc user.UseCase) transaction.UseCase {
	c := cache.NewTransactionCacheMap()
	rep := repository.NewTransactionRepositoryPSQL(pool, l)
	return usecase.NewTransactionUC(l, c, rep, uuc)
}
