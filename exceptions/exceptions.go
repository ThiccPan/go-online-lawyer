package exceptions

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrValidationFailed = errors.New("validation failed")
)