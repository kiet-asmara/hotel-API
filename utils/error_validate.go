package utils

import (
	"errors"

	"github.com/go-playground/validator"
)

var Validate *validator.Validate = validator.New()

func ErrorBind(err error) string {
	var ve validator.ValidationErrors
	out := ""
	if errors.As(err, &ve) {
		for _, fe := range ve {
			out = fe.Field() + ": " + msgForTag(fe.Tag())
		}
		return out
	}
	return out
}

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "alpha":
		return "must be alphabetical"
	case "gte":
		return "input to low"
	}
	return ""
}
