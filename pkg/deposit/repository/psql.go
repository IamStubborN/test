package repository

import (
	"github.com/jmoiron/sqlx"

	"github.com/IamStubborN/test/models"
	"github.com/IamStubborN/test/pkg/deposit"
	"github.com/IamStubborN/test/pkg/logger"
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

func (dr *depositRepository) GetAllDeposits() ([]*models.Deposit, error) {
	query := `select id, user_id, amount, balance_before, balance_after, date from deposits`

	var deposits []*models.Deposit
	err := dr.pool.Select(&deposits, query)
	if err != nil {
		return nil, err
	}

	return deposits, nil
}

func (dr *depositRepository) BackupDeposits(deposits []*models.Deposit) error {
	tx, err := dr.pool.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		switch err.(type) {
		case nil:
			dr.checkError(tx.Commit)
		default:
			dr.checkError(tx.Rollback)
		}
	}()

	query := `insert into deposits(
                     id, user_id, amount, balance_before, balance_after, date) 
                     values (:id, :user_id, :amount, :balance_before, :balance_after, :date)`

	stmt, err := tx.PrepareNamed(query)
	if err != nil {
		return err
	}

	for _, d := range deposits {
		_, err := stmt.Exec(d)
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
