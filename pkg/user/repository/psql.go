package repository

import (
	"github.com/jmoiron/sqlx"

	"github.com/IamStubborN/test/models"
	"github.com/IamStubborN/test/pkg/logger"
	"github.com/IamStubborN/test/pkg/user"
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

func (ur *userRepository) GetAllUsers() ([]*models.User, error) {
	query := `select id, balance from users`

	var users []*models.User
	err := ur.pool.Select(&users, query)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *userRepository) BackupUsers(users []*models.User) error {
	tx, err := ur.pool.Beginx()
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
	stmt, err := tx.Preparex(query)
	if err != nil {
		return err
	}

	for _, u := range users {
		_, err := stmt.Exec(u.ID, u.Balance)
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
