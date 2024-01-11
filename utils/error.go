package utils

import (
	"errors"
)

// service layer

type Error struct {
	svcErr error // generic http error
	appErr error // actual error
}

func (e Error) AppError() error {
	return e.appErr
}

func (e Error) ServiceError() error {
	return e.svcErr
}

func NewError(svcErr, appErr error) error {
	return Error{
		svcErr: svcErr,
		appErr: appErr,
	}
}

// so that Error = error
func (e Error) Error() string {
	return errors.Join(e.svcErr, e.appErr).Error()
}
