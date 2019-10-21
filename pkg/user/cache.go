package user

import "github.com/IamStubborN/test/models"

type Cache interface {
	AddUser(user *models.User)
	GetUser(userID uint64) *models.User
	ChangeUserBalance(userID uint64, balance float64)
	IsUserExist(userID uint64) bool
	GetBackupUsers() []*models.User
	CleanBackupUsers()
	PutUsersToCache(users []*models.User)
}
