package deposit

import "github.com/IamStubborN/test/models"

type Cache interface {
	AddDeposit(deposit *models.Deposit)
	IsDepositExist(depositID uint64) bool
	GetDepositCountAndSum(userID uint64) []float64
	GetBackupDeposits() []*models.Deposit
	CleanBackupDeposits()
	PutDepositsToCache(deposits []*models.Deposit)
}
