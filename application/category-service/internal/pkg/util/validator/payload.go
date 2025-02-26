// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package validator

import (
	"github.com/sirupsen/logrus"

	"category-service/internal/pkg/config"
)

type IPayload interface {
	CustomValidatePayload() error
}

func ValidatePayload(payload IPayload) error {
	validator := InitValidator()

	if err := validator.Validate(payload); err != nil {
		config.GetLogConfig().WithFields(logrus.Fields{
			"message": "validate payload fail",
			"detail":  err,
		}).Errorln("[VALIDATE PAYLOAD]")
		return err
	}
	if err := payload.CustomValidatePayload(); err != nil {
		return err
	}
	return nil
}
