package service

import (
	"context"
	"time"

	"github.com/IamStubborN/test/daemon"
	"github.com/IamStubborN/test/pkg/deposit"
	"github.com/IamStubborN/test/pkg/logger"
	"github.com/IamStubborN/test/pkg/transaction"
	"github.com/IamStubborN/test/pkg/user"
)

type back struct {
	freq        time.Duration
	logger      logger.Logger
	user        user.UseCase
	deposit     deposit.UseCase
	transaction transaction.UseCase
}

func NewBackupDaemon(
	freq time.Duration,
	l logger.Logger,
	uuc user.UseCase,
	duc deposit.UseCase,
	tuc transaction.UseCase) daemon.Daemon {
	return &back{
		freq:        freq,
		logger:      l,
		user:        uuc,
		deposit:     duc,
		transaction: tuc,
	}
}

func (b *back) Run(ctx context.Context) {
	if err := b.restore(); err != nil {
		b.logger.Warn(err)
	}

	for {
		select {
		case <-ctx.Done():
			b.logger.Info("backup: Server closed")
			return
		case <-time.NewTicker(b.freq).C:
			if err := b.backup(); err != nil {
				b.logger.Warn(err)
			}
		}
	}
}

func (b *back) backup() error {
	err := b.user.BackupUsers()
	if err != nil {
		return err
	}

	err = b.deposit.BackupDeposits()
	if err != nil {
		return err
	}

	err = b.transaction.BackupTransactions()
	if err != nil {
		return err
	}

	return nil
}

func (b *back) restore() error {
	err := b.user.RestoreUsers()
	if err != nil {
		return err
	}

	err = b.deposit.RestoreDeposits()
	if err != nil {
		return err
	}

	err = b.transaction.RestoreTransactions()
	if err != nil {
		return err
	}

	return nil
}
