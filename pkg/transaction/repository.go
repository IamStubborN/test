package transaction

import (
	"github.com/IamStubborN/test/models"
)

type Repository interface {
	GetAllTransactions() ([]*models.Transaction, error)
	BackupTransactions(transactions []*models.Transaction) error
}
