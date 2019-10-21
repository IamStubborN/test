package http

import "errors"

var (
	ErrEmptyBody = errors.New("user: empty body")
	ErrAddUser = errors.New("can't add user")
	ErrGetUser = errors.New("user: get")
	ErrUnmarshalBody = errors.New("user: decode body")
)
