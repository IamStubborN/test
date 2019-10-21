package usecase

import (
	"context"
	"github.com/IamStubborN/test/models"
	"github.com/IamStubborN/test/pkg/logger"
	"github.com/IamStubborN/test/pkg/user"
)

type userUC struct {
	logger     logger.Logger
	cache      user.Cache
	repository user.Repository
}

func NewUserUC(
	logger logger.Logger,
	cache user.Cache,
	repository user.Repository,
) user.UseCase {
	return &userUC{
		logger:     logger,
		cache:      cache,
		repository: repository,
	}
}

func (uuc *userUC) AddUser(user *models.User) error {
	if uuc.cache.IsUserExist(user.ID) {
		return ErrUserIsAlreadyExist
	}

	uuc.cache.AddUser(user)

	return nil
}

func (uuc *userUC) GetUser(userID uint64) (*models.User, error) {
	if !uuc.cache.IsUserExist(userID) {
		return nil, ErrUserIsNotExist
	}

	return uuc.cache.GetUser(userID), nil
}

func (uuc *userUC) ChangeUserBalance(userID uint64, balance float64) {
	uuc.cache.ChangeUserBalance(userID, balance)
}

func (uuc *userUC) IsUserExist(userID uint64) error {
	if !uuc.cache.IsUserExist(userID) {
		return ErrUserIsNotExist
	}

	return nil
}

func (uuc *userUC) RestoreUsers() error {
	users, err := uuc.repository.GetAllUsers(context.Background())
	if err != nil {
		return err
	}

	uuc.cache.PutUsersToCache(users)

	return nil
}

func (uuc *userUC) BackupUsers() error {
	users := uuc.cache.GetBackupUsers()
	err := uuc.repository.BackupUsers(context.Background(), users)
	if err != nil {
		return err
	}

	uuc.cache.CleanBackupUsers()
	return nil
}
