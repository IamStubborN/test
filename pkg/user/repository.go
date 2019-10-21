package user

import (
	"github.com/IamStubborN/test/models"
)

type Repository interface {
	GetAllUsers() ([]*models.User, error)
	BackupUsers(users []*models.User) error
}
