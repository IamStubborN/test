package instance

import "errors"

var (
	ErrEmptyBody    = errors.New("middleware: empty body")
	ErrInvalidBody  = errors.New("middleware: invalid body")
	ErrInvalidToken = errors.New("middleware: invalid token")
)
