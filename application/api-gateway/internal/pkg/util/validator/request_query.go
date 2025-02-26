// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package validator

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"api-gateway/internal/pkg/config"
)

type IRequestQuery interface {
	CustomValidateRequestQuery() error
}

func ValidateRequestQuery(ctx *fiber.Ctx, req IRequestQuery) error {
	validator := InitValidator()

	if err := ctx.QueryParser(req); err != nil {
		config.GetLogConfig().WithFields(logrus.Fields{
			"message": "bind request query fail",
			"detail":  err,
		}).Errorln("[VALIDATE REQUEST QUERY]")
		return err
	}
	if err := validator.Validate(req); err != nil {
		config.GetLogConfig().WithFields(logrus.Fields{
			"message": "validate request query fail",
			"detail":  err,
		}).Errorln("[VALIDATE REQUEST QUERY]")
		return err
	}
	if err := req.CustomValidateRequestQuery(); err != nil {
		return err
	}
	return nil
}
