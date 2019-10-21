package http

import "errors"

var (
	ErrEmptyBody = errors.New("transaction: response body is empty")
)
