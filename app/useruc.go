package app

import (
	"github.com/IamStubborN/test/pkg/logger"
	"github.com/IamStubborN/test/pkg/user"
	"github.com/IamStubborN/test/pkg/user/cache"
	"github.com/IamStubborN/test/pkg/user/repository"
	"github.com/IamStubborN/test/pkg/user/usecase"
	"github.com/jmoiron/sqlx"
)

func initializeUserUC(pool *sqlx.DB, l logger.Logger) user.UseCase {
	c := cache.NewUsersCacheMap()
	rep := repository.NewUserRepositoryPSQL(pool, l)
	return usecase.NewUserUC(l, c, rep)
}
