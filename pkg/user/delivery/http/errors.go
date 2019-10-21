package http

import "errors"

var (
	ErrEmptyBody     = errors.New("user: empty body")
	ErrGetUser       = errors.New("user: get")
	ErrUnmarshalBody = errors.New("user: decode body")
)
