// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package validator

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"api-gateway/internal/pkg/config"
)

type IRequestBody interface {
	CustomValidateRequestBody() error
}

func ValidateRequestBody(ctx *fiber.Ctx, req IRequestBody) error {
	validator := InitValidator()

	if err := ctx.BodyParser(req); err != nil {
		config.GetLogConfig().WithFields(logrus.Fields{
			"message": "bind request body fail",
			"detail":  err,
		}).Errorln("[VALIDATE REQUEST BODY]")
		return err
	}
	if err := validator.Validate(req); err != nil {
		config.GetLogConfig().WithFields(logrus.Fields{
			"message": "validate request body fail",
			"detail":  err,
		}).Errorln("[VALIDATE REQUEST BODY]")
		return err
	}
	if err := req.CustomValidateRequestBody(); err != nil {
		return err
	}
	return nil
}
