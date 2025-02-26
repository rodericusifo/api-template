// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package rest

import (
	"github.com/gofiber/fiber/v2"

	"api-gateway/internal/domain/core/category/controller/http/rest/handler"
	"api-gateway/internal/pkg/config"

	jwtware "github.com/gofiber/contrib/jwt"

	external_client_authservice_http_grpc_handler_user "api-gateway/external/client/auth-service/http/grpc/handler/user"
	external_client_bookservice_http_grpc_handler_book "api-gateway/external/client/book-service/http/grpc/handler/book"
	external_client_categoryservice_http_grpc_handler_category "api-gateway/external/client/category-service/http/grpc/handler/category"
)

func InitREST(
	router fiber.Router,
	userClientHandler external_client_authservice_http_grpc_handler_user.IUserClientHandler,
	categoryClientHandler external_client_categoryservice_http_grpc_handler_category.ICategoryClientHandler,
	bookClientHandler external_client_bookservice_http_grpc_handler_book.IBookClientHandler,
) {
	category := router.Group("/categories")
	category.Use(jwtware.New(config.GetAuthConfig()))
	categoryHandler := handler.InitCategoryHandler(userClientHandler, categoryClientHandler, bookClientHandler)
	categoryHandler.Mount(category)
}
