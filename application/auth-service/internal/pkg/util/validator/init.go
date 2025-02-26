// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package validator

import (
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	Validator *validator.Validate
}

func InitValidator() *Validator {
	v := validator.New()

	return &Validator{Validator: v}
}

func (cv *Validator) Validate(i any) error {
	if err := cv.Validator.Struct(i); err != nil {
		return err
	}
	return nil
}
