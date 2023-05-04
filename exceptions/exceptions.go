package exceptions

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)