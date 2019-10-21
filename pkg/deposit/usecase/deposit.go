package usecase

import (
	"context"
	"time"

	"github.com/shopspring/decimal"

	"github.com/IamStubborN/test/models"
	"github.com/IamStubborN/test/pkg/deposit"
	"github.com/IamStubborN/test/pkg/logger"
	"github.com/IamStubborN/test/pkg/user"
)

type depositUC struct {
	logger     logger.Logger
	cache      deposit.Cache
	repository deposit.Repository
	user       user.UseCase
}

func NewDepositUC(
	logger logger.Logger,
	cache deposit.Cache,
	repository deposit.Repository,
	user user.UseCase,
) deposit.UseCase {
	return &depositUC{
		logger:     logger,
		cache:      cache,
		repository: repository,
		user:       user,
	}
}

func (duc *depositUC) AddDeposit(deposit *models.Deposit) error {
	u, err := duc.user.GetUser(deposit.UserID)
	if err != nil {
		return err
	}

	if duc.cache.IsDepositExist(deposit.ID) {
		return ErrDepositAlreadyExist
	}

	if deposit.Amount <= 0 {
		return ErrDepositAmount
	}

	deposit.BalanceBefore = u.Balance

	before := decimal.NewFromFloat(deposit.BalanceBefore)
	amount := decimal.NewFromFloat(deposit.Amount)

	result, _ := before.Add(amount).Float64()
	deposit.BalanceAfter = result
	deposit.Date = time.Now().UTC()

	duc.cache.AddDeposit(deposit)
	duc.user.ChangeUserBalance(deposit.UserID, deposit.BalanceAfter)

	return nil
}

func (duc *depositUC) GetDepositCountAndSum(userID uint64) (count uint64, sum float64) {
	depositValues := duc.cache.GetDepositCountAndSum(userID)

	var result decimal.Decimal
	for _, value := range depositValues {
		result = result.Add(decimal.NewFromFloat(value))
	}

	count = uint64(len(depositValues))
	sum, _ = result.Float64()

	return
}

func (duc *depositUC) RestoreDeposits() error {
	deposits, err := duc.repository.GetAllDeposits(context.Background())
	if err != nil {
		return err
	}

	if len(deposits) > 0 {
		duc.cache.PutDepositsToCache(deposits)
	}

	return nil
}

func (duc *depositUC) BackupDeposits() error {
	deposits := duc.cache.GetBackupDeposits()

	if len(deposits) == 0 {
		return nil
	}

	err := duc.repository.BackupDeposits(context.Background(), deposits)
	if err != nil {
		return err
	}

	duc.cache.CleanBackupDeposits()
	return nil
}
