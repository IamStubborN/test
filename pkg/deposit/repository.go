package deposit

import (
	"context"

	"github.com/IamStubborN/test/models"
)

type Repository interface {
	GetAllDeposits(ctx context.Context) ([]*models.Deposit, error)
	BackupDeposits(ctx context.Context, deposits []*models.Deposit) error
}
