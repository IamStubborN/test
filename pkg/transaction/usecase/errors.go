package usecase

import "errors"

var (
	ErrTransactionAlreadyExist = errors.New("transaction: transaction is already exist")
	ErrTransactionAmount       = errors.New("transaction: amount is invalid")
	ErrNotEnoughMoney          = errors.New("transaction: not enough money")
)
