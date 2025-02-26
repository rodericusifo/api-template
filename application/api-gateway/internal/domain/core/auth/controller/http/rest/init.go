// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package rest

import (
	"github.com/gofiber/fiber/v2"

	"api-gateway/internal/domain/core/auth/controller/http/rest/handler"

	external_client_authservice_http_grpc_handler_auth "api-gateway/external/client/auth-service/http/grpc/handler/auth"
)

func InitREST(
	router fiber.Router,
	authClientHandler external_client_authservice_http_grpc_handler_auth.IAuthClientHandler,
) {
	auth := router.Group("/auth")
	authHandler := handler.InitAuthHandler(authClientHandler)
	authHandler.Mount(auth)
}
