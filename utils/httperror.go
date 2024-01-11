package utils

import (
	"errors"
	"net/http"
)

type APIError struct {
	Status  int
	Message string
}

func FromError(err error) APIError {
	var svcError Error
	var apiError APIError

	// err = from newerror
	// errors as checks if both are type utils.Error
	if errors.As(err, &svcError) {
		// set actual error on message
		apiError.Message = svcError.AppError().Error()
		// check error
		svcErr := svcError.ServiceError()
		switch svcErr {
		case ErrFailedBind:
			apiError.Message = ErrorBind(svcError.AppError()) // check which field fails validation
			apiError.Status = http.StatusBadRequest
		case ErrBadRequest:
			apiError.Status = http.StatusBadRequest
		case ErrInternalFailure:
			apiError.Status = http.StatusInternalServerError
		case ErrNotFound:
			apiError.Status = http.StatusNotFound
		case ErrUnauthorized:
			apiError.Status = http.StatusUnauthorized
		}
	}

	return apiError
}
