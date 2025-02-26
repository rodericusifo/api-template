// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"api-gateway/internal/domain/core/auth/controller/http/rest/request"
	"api-gateway/internal/pkg/util/response"
	"api-gateway/internal/pkg/util/validator"

	pb_auth "api-gateway/external/client/auth-service/http/grpc/proto/auth"
)

func (h *AuthHandler) RegisterAuth(ctx *fiber.Ctx) error {
	reqBody := new(request.RegisterAuthRequestBody)
	if err := validator.ValidateRequestBody(ctx, reqBody); err != nil {
		return err
	}

	err := h.AuthClientHandler.RegisterAuth(context.Background(), &pb_auth.RegisterAuthRequest{
		Name:     reqBody.Name,
		Email:    reqBody.Email,
		Password: reqBody.Password,
		RoleSlug: reqBody.RoleSlug,
	})
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response.ResponseSuccess[any]("auth register success", nil, nil))
}
