package app

import (
	"github.com/IamStubborN/test/config"
	"github.com/IamStubborN/test/pkg/backup/service"
	"github.com/IamStubborN/test/pkg/deposit"
	"github.com/IamStubborN/test/pkg/logger"
	"github.com/IamStubborN/test/pkg/transaction"
	"github.com/IamStubborN/test/pkg/user"
	"github.com/IamStubborN/test/worker"
)

func initializeBackupDaemon(
	config *config.Config,
	l logger.Logger,
	uuc user.UseCase,
	duc deposit.UseCase,
	tuc transaction.UseCase) worker.Worker {

	return service.NewBackupDaemon(config.Cache.BackupFreq, l, uuc, duc, tuc)
}
