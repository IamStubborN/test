package deposit

import (
	"github.com/IamStubborN/test/models"
)

type Repository interface {
	GetAllDeposits() ([]*models.Deposit, error)
	BackupDeposits(deposits []*models.Deposit) error
}
