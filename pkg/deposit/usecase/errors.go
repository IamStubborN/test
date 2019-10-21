package usecase

import "errors"

var (
	ErrDepositAlreadyExist = errors.New("deposit: deposit is already exist")
	ErrDepositAmount       = errors.New("deposit: amount is invalid")
)
