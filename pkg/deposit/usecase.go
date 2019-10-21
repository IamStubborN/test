package deposit

import "github.com/IamStubborN/test/models"

type UseCase interface {
	AddDeposit(deposit *models.Deposit) error
	GetDepositCountAndSum(userID uint64) (count uint64, sum float64)
	BackupDeposits() error
	RestoreDeposits() error
}
