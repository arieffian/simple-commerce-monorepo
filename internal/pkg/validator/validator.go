package validator

import (
	go_validator "github.com/go-playground/validator/v10"
)

type ValidatorService interface {
	Validate(interface{}) error
}

type validatorService struct {
	validator *go_validator.Validate
}

func NewValidatorService() ValidatorService {
	v := go_validator.New()

	return &validatorService{
		validator: v,
	}
}

func (v *validatorService) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
