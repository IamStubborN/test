package transaction

import (
	"context"

	"github.com/IamStubborN/test/models"
)

type Repository interface {
	GetAllTransactions(ctx context.Context) ([]*models.Transaction, error)
	BackupTransactions(ctx context.Context, transactions []*models.Transaction) error
}
