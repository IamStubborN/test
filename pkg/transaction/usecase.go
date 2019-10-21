package transaction

import "github.com/IamStubborN/test/models"

type UseCase interface {
	AddTransaction(transaction *models.Transaction) error
	GetWinCountAndSum(userID uint64) (count uint64, sum float64)
	GetBetCountAndSum(userID uint64) (count uint64, sum float64)
	BackupTransactions() error
	RestoreTransactions() error
}
