package app

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/IamStubborN/test/pkg/logger"

	"github.com/IamStubborN/test/config"
	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
)

func initializeSQLConn(cfg *config.Config, logger logger.Logger) *sqlx.DB {
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Host, strconv.Itoa(cfg.DB.Port), cfg.DB.User, cfg.DB.Password, cfg.DB.Name)
	pool, err := sqlx.Open("pgx", dbInfo)
	if err != nil {
		logger.Fatal(err)
		return nil
	}

	if err := retryConnect(pool, cfg.DB.RetryCount, logger); err != nil {
		logger.Fatal(err)
	}

	migrationLogic(pool, logger)

	return pool
}

func retryConnect(pool *sqlx.DB, fatalRetry int, logger logger.Logger) error {
	var retryCount int
	for range time.NewTicker(time.Second).C {
		if fatalRetry == retryCount {
			return errors.New("can't connect to database")
		}

		retryCount++
		if err := pool.Ping(); err != nil {
			logger.WithFields([]interface{}{
				"status", "retrying",
				"try", retryCount,
			}).Info("connect to db")
			continue
		}

		logger.WithFields([]interface{}{
			"status", "connected",
		}).Info("connect to db")
		break
	}

	return nil
}

func migrationLogic(db *sqlx.DB, logger logger.Logger) {
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	_, err := migrate.Exec(db.DB, "postgres", migrations, migrate.Up)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("migrations complete")
}
