package validator

import (
	"github.com/go-playground/validator/v10"
)

type GlobalValidator struct {
	Validator *validator.Validate
}
