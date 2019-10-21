package repository

import (
	"context"
	"database/sql"

	"github.com/IamStubborN/test/models"
	"github.com/IamStubborN/test/pkg/deposit"
	"github.com/IamStubborN/test/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type depositRepository struct {
	pool   *sqlx.DB
	logger logger.Logger
}

func NewDepositRepositoryPSQL(pool *sqlx.DB, l logger.Logger) deposit.Repository {
	return &depositRepository{
		pool:   pool,
		logger: l,
	}
}

func (dr *depositRepository) GetAllDeposits(ctx context.Context) ([]*models.Deposit, error) {
	query := `select id, user_id, amount, balance_before, balance_after, date from deposits`

	var deposits []*models.Deposit
	err := dr.pool.Select(&deposits, query)
	if err != nil {
		return nil, err
	}

	return deposits, nil
}

func (dr *depositRepository) BackupDeposits(ctx context.Context, deposits []*models.Deposit) error {
	tx, err := dr.pool.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	defer func() {
		switch err.(type) {
		case nil:
			dr.checkError(tx.Commit)
		default:
			dr.checkError(tx.Rollback)
			return
		}

	}()

	query := `insert into deposits(
                     id, user_id, amount, balance_before, balance_after, date) 
                     values (:id, :user_id, :amount, :balance_before, :balance_after, :date)`

	stmt, err := tx.PrepareNamedContext(ctx, query)
	if err != nil {
		return err
	}

	for _, d := range deposits {
		_, err := stmt.ExecContext(ctx, &d)
		if err != nil {
			return err
		}
	}

	return nil
}

func (dr *depositRepository) checkError(fn func() error) {
	if err := fn(); err != nil {
		dr.logger.Warn(err)
	}
}
