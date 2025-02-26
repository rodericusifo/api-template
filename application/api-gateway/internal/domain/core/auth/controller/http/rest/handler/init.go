// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"github.com/gofiber/fiber/v2"

	external_client_authservice_http_grpc_handler_auth "api-gateway/external/client/auth-service/http/grpc/handler/auth"
)

type AuthHandler struct {
	AuthClientHandler external_client_authservice_http_grpc_handler_auth.IAuthClientHandler
}

func InitAuthHandler(
	authClientHandler external_client_authservice_http_grpc_handler_auth.IAuthClientHandler,
) *AuthHandler {
	return &AuthHandler{
		AuthClientHandler: authClientHandler,
	}
}

func (authHandler *AuthHandler) Mount(group fiber.Router) {
	group.Post("/register", authHandler.RegisterAuth)
	group.Post("/login", authHandler.LoginAuth)
}
