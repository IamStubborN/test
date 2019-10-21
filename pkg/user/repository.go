package user

import (
	"context"

	"github.com/IamStubborN/test/models"
)

type Repository interface {
	GetAllUsers(ctx context.Context) ([]*models.User, error)
	BackupUsers(ctx context.Context, users []*models.User) error
}
