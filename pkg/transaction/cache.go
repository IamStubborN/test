package transaction

import "github.com/IamStubborN/test/models"

type Cache interface {
	AddTransaction(transaction *models.Transaction)
	IsTransactionExist(transactionID uint64) bool
	GetWinCountAndSum(userID uint64) []float64
	GetBetCountAndSum(userID uint64) []float64
	GetBackupTransactions() []*models.Transaction
	CleanBackupTransactions()
	PutTransactionsToCache(transactions []*models.Transaction)
}
