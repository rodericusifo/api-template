// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package validator

import (
	"github.com/sirupsen/logrus"

	"api-gateway/internal/pkg/config"
)

type IRequestUser interface {
	CustomValidateRequestUser() error
}

func ValidateRequestUser(req IRequestUser) error {
	validator := InitValidator()

	if err := validator.Validate(req); err != nil {
		config.GetLogConfig().WithFields(logrus.Fields{
			"message": "validate request user fail",
			"detail":  err,
		}).Errorln("[VALIDATE REQUEST USER]")
		return err
	}
	if err := req.CustomValidateRequestUser(); err != nil {
		return err
	}
	return nil
}
