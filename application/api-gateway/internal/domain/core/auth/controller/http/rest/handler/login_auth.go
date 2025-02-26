// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"api-gateway/internal/domain/core/auth/controller/http/rest/request"
	"api-gateway/internal/domain/core/auth/controller/http/rest/response"
	"api-gateway/internal/pkg/util/validator"

	pb_auth "api-gateway/external/client/auth-service/http/grpc/proto/auth"
	internal_pkg_util_response "api-gateway/internal/pkg/util/response"
)

func (h *AuthHandler) LoginAuth(ctx *fiber.Ctx) error {
	reqBody := new(request.LoginAuthRequestBody)
	if err := validator.ValidateRequestBody(ctx, reqBody); err != nil {
		return err
	}

	authLoginRes, err := h.AuthClientHandler.LoginAuth(context.Background(), &pb_auth.LoginAuthRequest{
		Email:    reqBody.Email,
		Password: reqBody.Password,
	})
	if err != nil {
		return err
	}

	loginAuthRes := &response.LoginAuthResponse{
		Token: authLoginRes.GetToken(),
	}

	return ctx.Status(fiber.StatusOK).JSON(internal_pkg_util_response.ResponseSuccess("auth login success", loginAuthRes, nil))
}
