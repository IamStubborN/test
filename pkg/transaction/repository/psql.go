package repository

import (
	"context"
	"database/sql"

	"github.com/IamStubborN/test/models"
	"github.com/IamStubborN/test/pkg/logger"
	"github.com/IamStubborN/test/pkg/transaction"
	"github.com/jmoiron/sqlx"
)

type transactionRepository struct {
	pool   *sqlx.DB
	logger logger.Logger
}

func NewTransactionRepositoryPSQL(pool *sqlx.DB, l logger.Logger) transaction.Repository {
	return &transactionRepository{
		pool:   pool,
		logger: l,
	}
}

func (tr *transactionRepository) GetAllTransactions(ctx context.Context) ([]*models.Transaction, error) {
	query := `select id, user_id, amount, tt.type_id, tt.name, balance_before, balance_after, date from transactions as t 
				inner join transaction_types tt on t.transaction_type_id = tt.type_id`

	var transactions []*models.Transaction
	err := tr.pool.Select(&transactions, query)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (tr *transactionRepository) BackupTransactions(ctx context.Context, transactions []*models.Transaction) error {
	tx, err := tr.pool.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	defer func() {
		switch err.(type) {
		case nil:
			tr.checkError(tx.Commit)
		default:
			tr.checkError(tx.Rollback)
			return
		}

	}()
	var transactionTypes []*models.TransactionType
	getTransactionTypesQuery := `select type_id, name  from transaction_types`
	err = tx.Select(&transactionTypes, getTransactionTypesQuery)
	if err != nil {
		return err
	}

	insertTransactionQuery := `insert into transactions(id, user_id, transaction_type_id, amount,  balance_after, balance_before, date) 
                     values (:id, :user_id, :transaction_type_id, :amount, :balance_before, :balance_after, :date)`

	stmt, err := tx.PrepareNamedContext(ctx, insertTransactionQuery)
	if err != nil {
		return err
	}

	for _, t := range transactions {

		t.TransactionType.ID = tr.getTransactionTypeID(t.Name, transactionTypes)

		argQ := map[string]interface{}{
			"id":                  t.ID,
			"user_id":             t.UserID,
			"amount":              t.Amount,
			"transaction_type_id": t.TransactionType.ID,
			"balance_before":      t.BalanceBefore,
			"balance_after":       t.BalanceAfter,
			"date":                t.Date,
		}

		_, err := stmt.ExecContext(ctx, argQ)
		if err != nil {
			return err
		}
	}

	return nil
}

func (tr *transactionRepository) getTransactionTypeID(name string, transactionTypes []*models.TransactionType) uint8 {
	for _, transactionType := range transactionTypes {
		if transactionType.Name == name {
			return transactionType.ID
		}
	}

	return 0
}

func (tr *transactionRepository) checkError(fn func() error) {
	if err := fn(); err != nil {
		tr.logger.Warn(err)
	}
}
