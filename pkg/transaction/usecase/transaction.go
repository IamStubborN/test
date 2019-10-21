package usecase

import (
	"time"

	"github.com/shopspring/decimal"

	"github.com/IamStubborN/test/models"
	"github.com/IamStubborN/test/pkg/logger"
	"github.com/IamStubborN/test/pkg/transaction"
	"github.com/IamStubborN/test/pkg/user"
)

type transactionUC struct {
	logger     logger.Logger
	cache      transaction.Cache
	repository transaction.Repository
	user       user.UseCase
}

func NewTransactionUC(
	logger logger.Logger,
	cache transaction.Cache,
	repository transaction.Repository,
	user user.UseCase,
) transaction.UseCase {
	return &transactionUC{
		logger:     logger,
		cache:      cache,
		repository: repository,
		user:       user,
	}
}

func (tuc *transactionUC) AddTransaction(transaction *models.Transaction) error {
	u, err := tuc.user.GetUser(transaction.UserID)
	if err != nil {
		return err
	}

	if tuc.cache.IsTransactionExist(transaction.ID) {
		return ErrTransactionAlreadyExist
	}

	if transaction.Amount <= 0 {
		return ErrTransactionAmount
	}

	transaction.BalanceBefore = u.Balance

	before := decimal.NewFromFloat(transaction.BalanceBefore)
	amount := decimal.NewFromFloat(transaction.Amount)

	switch transaction.Name {
	case "Win":
		result, _ := before.Add(amount).Float64()
		transaction.BalanceAfter = result
	case "Bet":
		if u.Balance < transaction.Amount {
			return ErrNotEnoughMoney
		}

		result, _ := before.Sub(amount).Float64()
		transaction.BalanceAfter = result
	}

	transaction.Date = time.Now().UTC()

	tuc.cache.AddTransaction(transaction)
	tuc.user.ChangeUserBalance(transaction.UserID, transaction.BalanceAfter)

	return nil
}

func (tuc *transactionUC) GetWinCountAndSum(userID uint64) (count uint64, sum float64) {
	winValues := tuc.cache.GetWinCountAndSum(userID)

	var result decimal.Decimal
	for _, value := range winValues {
		result = result.Add(decimal.NewFromFloat(value))
	}

	count = uint64(len(winValues))
	sum, _ = result.Float64()

	return
}

func (tuc *transactionUC) GetBetCountAndSum(userID uint64) (count uint64, sum float64) {
	betValues := tuc.cache.GetBetCountAndSum(userID)

	var result decimal.Decimal
	for _, value := range betValues {
		result = result.Add(decimal.NewFromFloat(value))
	}

	count = uint64(len(betValues))
	sum, _ = result.Float64()

	return
}

func (tuc *transactionUC) RestoreTransactions() error {
	transactions, err := tuc.repository.GetAllTransactions()
	if err != nil {
		return err
	}

	tuc.cache.PutTransactionsToCache(transactions)

	return nil
}

func (tuc *transactionUC) BackupTransactions() error {
	transactions := tuc.cache.GetBackupTransactions()

	if len(transactions) == 0 {
		return nil
	}

	err := tuc.repository.BackupTransactions(transactions)
	if err != nil {
		return err
	}

	tuc.cache.CleanBackupTransactions()
	return nil
}
