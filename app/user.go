package app

import (
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"

	"github.com/IamStubborN/test/pkg/deposit"
	"github.com/IamStubborN/test/pkg/logger"
	"github.com/IamStubborN/test/pkg/responder"
	"github.com/IamStubborN/test/pkg/transaction"
	"github.com/IamStubborN/test/pkg/user"
	"github.com/IamStubborN/test/pkg/user/cache"
	"github.com/IamStubborN/test/pkg/user/delivery/http"
	"github.com/IamStubborN/test/pkg/user/repository"
	"github.com/IamStubborN/test/pkg/user/usecase"
)

func registerUserHandlers(
	router chi.Router,
	l logger.Logger,
	r responder.Responder,
	uuc user.UseCase,
	duc deposit.UseCase,
	tuc transaction.UseCase) {
	http.RegisterUserHandler(router, l, r, uuc, duc, tuc)
}

func initializeUserUC(pool *sqlx.DB, l logger.Logger) user.UseCase {
	c := cache.NewUsersCacheMap()
	rep := repository.NewUserRepositoryPSQL(pool, l)
	return usecase.NewUserUC(l, c, rep)
}
