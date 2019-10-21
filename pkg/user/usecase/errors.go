package usecase

import "errors"

var (
	ErrUserIsAlreadyExist = errors.New("user: user is already exist")
	ErrUserIsNotExist     = errors.New("user: user isn't exist")
)
