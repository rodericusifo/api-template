// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package middleware

import (
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"api-gateway/internal/pkg/config"
)

func APIRecover(ctx *fiber.Ctx, err any) {
	config.GetLogConfig().WithFields(logrus.Fields{
		"message": "recovered from panic",
		"detail":  err,
	}).Errorln("[UNARY GRPC RECOVER]")
	config.GetLogConfig().Errorf("%s", debug.Stack())
}
