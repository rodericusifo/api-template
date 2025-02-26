// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"github.com/gofiber/fiber/v2"

	"api-gateway/internal/pkg/middleware"

	external_client_authservice_http_grpc_handler_user "api-gateway/external/client/auth-service/http/grpc/handler/user"
	external_client_authorservice_http_grpc_handler_author "api-gateway/external/client/author-service/http/grpc/handler/author"
	external_client_bookservice_http_grpc_handler_book "api-gateway/external/client/book-service/http/grpc/handler/book"
)

type AuthorHandler struct {
	UserClientHandler   external_client_authservice_http_grpc_handler_user.IUserClientHandler
	AuthorClientHandler external_client_authorservice_http_grpc_handler_author.IAuthorClientHandler
	BookClientHandler   external_client_bookservice_http_grpc_handler_book.IBookClientHandler
}

func InitAuthorHandler(
	userClientHandler external_client_authservice_http_grpc_handler_user.IUserClientHandler,
	authorClientHandler external_client_authorservice_http_grpc_handler_author.IAuthorClientHandler,
	bookClientHandler external_client_bookservice_http_grpc_handler_book.IBookClientHandler,
) *AuthorHandler {
	return &AuthorHandler{
		UserClientHandler:   userClientHandler,
		AuthorClientHandler: authorClientHandler,
		BookClientHandler:   bookClientHandler,
	}
}

func (authorHandler *AuthorHandler) Mount(group fiber.Router) {
	group.Post("/create", middleware.APIUser(authorHandler.UserClientHandler), middleware.APIUserRolePermissions(authorHandler.UserClientHandler), authorHandler.CreateAuthor)
	group.Get("/list", middleware.APIUser(authorHandler.UserClientHandler), middleware.APIUserRolePermissions(authorHandler.UserClientHandler), authorHandler.GetAuthors)
	group.Get("/:xid/detail", middleware.APIUser(authorHandler.UserClientHandler), middleware.APIUserRolePermissions(authorHandler.UserClientHandler), authorHandler.GetAuthor)
	group.Put("/:xid/update", middleware.APIUser(authorHandler.UserClientHandler), middleware.APIUserRolePermissions(authorHandler.UserClientHandler), authorHandler.UpdateAuthor)
	group.Delete("/:xid/delete", middleware.APIUser(authorHandler.UserClientHandler), middleware.APIUserRolePermissions(authorHandler.UserClientHandler), authorHandler.DeleteAuthor)
}
