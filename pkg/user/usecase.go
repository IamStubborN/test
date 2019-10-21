package user

import (
	"github.com/IamStubborN/test/models"
)

type UseCase interface {
	AddUser(user *models.User) error
	GetUser(userID uint64) (*models.User, error)
	ChangeUserBalance(userID uint64, balance float64)
	IsUserExist(userID uint64) error
	BackupUsers() error
	RestoreUsers() error
}
