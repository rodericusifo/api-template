// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package middleware

import (
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"api-gateway/internal/pkg/config"
	"api-gateway/internal/pkg/util/response"
)

func RESTError(ctx *fiber.Ctx, err error) error {
	config.GetLogConfig().WithFields(logrus.Fields{
		"type":   fmt.Sprintf("%T", err),
		"detail": err,
	}).Errorln("[API ERROR]")
	fe, ok := err.(*fiber.Error)
	if ok {
		return ctx.Status(fe.Code).JSON(response.ResponseFail(fe.Error(), fe))
	}
	ve, ok := err.(validator.ValidationErrors)
	if ok {
		type ErrorResponse struct {
			FailedField string `json:"failed_field"`
			Tag         string `json:"tag"`
			Error       string `json:"error"`
		}
		var errors []*ErrorResponse
		for _, err := range ve {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Error = err.Error()
			errors = append(errors, &element)
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ResponseFail("validation error", errors))
	}
	me, ok := err.(*json.MarshalerError)
	if ok {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(response.ResponseFail(me.Error(), me.Unwrap()))
	}
	re, ok := err.(runtime.Error)
	if ok {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(response.ResponseFail(re.Error(), re))
	}
	if e := err; e != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ResponseFail(e.Error(), e))
	}
	return nil
}
