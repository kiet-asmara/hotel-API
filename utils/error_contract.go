package utils

import "errors"

var (
	ErrBadRequest      = errors.New("bad request")
	ErrInternalFailure = errors.New("internal failure")
	ErrNotFound        = errors.New("not found")
	ErrFailedBind      = errors.New("failed bind json")
	ErrUnauthorized    = errors.New("access unauthorized")
)
