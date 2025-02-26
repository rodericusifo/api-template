// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package rest

import (
	"github.com/gofiber/fiber/v2"

	"api-gateway/internal/domain/core/author/controller/http/rest/handler"
	"api-gateway/internal/pkg/config"

	jwtware "github.com/gofiber/contrib/jwt"

	external_client_authservice_http_grpc_handler_user "api-gateway/external/client/auth-service/http/grpc/handler/user"
	external_client_authorservice_http_grpc_handler_author "api-gateway/external/client/author-service/http/grpc/handler/author"
	external_client_bookservice_http_grpc_handler_book "api-gateway/external/client/book-service/http/grpc/handler/book"
)

func InitREST(
	router fiber.Router,
	userClientHandler external_client_authservice_http_grpc_handler_user.IUserClientHandler,
	authorClientHandler external_client_authorservice_http_grpc_handler_author.IAuthorClientHandler,
	bookClientHandler external_client_bookservice_http_grpc_handler_book.IBookClientHandler,
) {
	author := router.Group("/authors")
	author.Use(jwtware.New(config.GetAuthConfig()))
	authorHandler := handler.InitAuthorHandler(userClientHandler, authorClientHandler, bookClientHandler)
	authorHandler.Mount(author)
}
