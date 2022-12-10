package errors

import "errors"

var (
	ErrTypeMismatch      = errors.New("received invalid type")
	ErrInvalidWorkerType = errors.New("invalid worker type string")
)
