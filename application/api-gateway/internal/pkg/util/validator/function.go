// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func isbn(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^[0-9\-]+$`)
	return re.MatchString(fl.Field().String())
}
