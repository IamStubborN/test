package repository

import (
	"context"
	"database/sql"

	"github.com/IamStubborN/test/models"
	"github.com/IamStubborN/test/pkg/logger"
	"github.com/IamStubborN/test/pkg/user"
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	pool   *sqlx.DB
	logger logger.Logger
}

func NewUserRepositoryPSQL(pool *sqlx.DB, l logger.Logger) user.Repository {
	return &userRepository{
		pool:   pool,
		logger: l,
	}
}

func (ur *userRepository) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	query := `select id, balance from users`

	var users []*models.User
	err := ur.pool.Select(&users, query)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *userRepository) BackupUsers(ctx context.Context, users []*models.User) error {
	tx, err := ur.pool.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	defer func() {
		switch err.(type) {
		case nil:
			ur.checkError(tx.Commit)
		default:
			ur.checkError(tx.Rollback)
		}

	}()

	query := `insert into users(id, balance) VALUES ($1, $2) 
		ON CONFLICT (id) DO UPDATE SET balance=$2`
	stmt, err := tx.PreparexContext(ctx, query)
	if err != nil {
		return err
	}

	for _, u := range users {
		_, err := stmt.ExecContext(ctx, u.ID, u.Balance)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ur *userRepository) checkError(fn func() error) {
	if err := fn(); err != nil {
		ur.logger.Warn(err)
	}
}
