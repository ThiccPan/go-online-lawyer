package exceptions

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrValidationFailed = errors.New("validation failed")
	ErrUserAlreadyExist = errors.New("user already exist")
	ErrInvalidTimeFormat = errors.New("invalid time format, use RFC3339 format")
	ErrKonsultasiNotFound = errors.New("konsultasi not found")
)