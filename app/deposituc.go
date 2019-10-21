package app

import (
	"github.com/jmoiron/sqlx"

	"github.com/IamStubborN/test/pkg/deposit"
	"github.com/IamStubborN/test/pkg/deposit/cache"
	"github.com/IamStubborN/test/pkg/deposit/repository"
	"github.com/IamStubborN/test/pkg/deposit/usecase"
	"github.com/IamStubborN/test/pkg/logger"
	"github.com/IamStubborN/test/pkg/user"
)

func initializeDepositUC(pool *sqlx.DB, l logger.Logger, uuc user.UseCase) deposit.UseCase {
	c := cache.NewDepositCacheMap()
	rep := repository.NewDepositRepositoryPSQL(pool, l)
	return usecase.NewDepositUC(l, c, rep, uuc)
}
