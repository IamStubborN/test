package app

import (
	"github.com/IamStubborN/test/pkg/logger"
	"github.com/IamStubborN/test/pkg/transaction"
	"github.com/IamStubborN/test/pkg/transaction/cache"
	"github.com/IamStubborN/test/pkg/transaction/repository"
	"github.com/IamStubborN/test/pkg/transaction/usecase"
	"github.com/IamStubborN/test/pkg/user"
	"github.com/jmoiron/sqlx"
)

func initializeTransactionUC(pool *sqlx.DB, l logger.Logger, uuc user.UseCase) transaction.UseCase {
	c := cache.NewTransactionCacheMap()
	rep := repository.NewTransactionRepositoryPSQL(pool, l)
	return usecase.NewTransactionUC(l, c, rep, uuc)
}

