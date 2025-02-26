// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"github.com/gofiber/fiber/v2"

	"api-gateway/internal/pkg/middleware"

	external_client_authservice_http_grpc_handler_user "api-gateway/external/client/auth-service/http/grpc/handler/user"
	external_client_authorservice_http_grpc_handler_author "api-gateway/external/client/author-service/http/grpc/handler/author"
	external_client_bookservice_http_grpc_handler_book "api-gateway/external/client/book-service/http/grpc/handler/book"
	external_client_categoryservice_http_grpc_handler_category "api-gateway/external/client/category-service/http/grpc/handler/category"
)

type BookHandler struct {
	UserClientHandler     external_client_authservice_http_grpc_handler_user.IUserClientHandler
	BookClientHandler     external_client_bookservice_http_grpc_handler_book.IBookClientHandler
	AuthorClientHandler   external_client_authorservice_http_grpc_handler_author.IAuthorClientHandler
	CategoryClientHandler external_client_categoryservice_http_grpc_handler_category.ICategoryClientHandler
}

func InitBookHandler(
	userClientHandler external_client_authservice_http_grpc_handler_user.IUserClientHandler,
	bookClientHandler external_client_bookservice_http_grpc_handler_book.IBookClientHandler,
	authorClientHandler external_client_authorservice_http_grpc_handler_author.IAuthorClientHandler,
	categoryClientHandler external_client_categoryservice_http_grpc_handler_category.ICategoryClientHandler,
) *BookHandler {
	return &BookHandler{
		UserClientHandler:     userClientHandler,
		BookClientHandler:     bookClientHandler,
		AuthorClientHandler:   authorClientHandler,
		CategoryClientHandler: categoryClientHandler,
	}
}

func (bookHandler *BookHandler) Mount(group fiber.Router) {
	group.Post("/create", middleware.APIUser(bookHandler.UserClientHandler), middleware.APIUserRolePermissions(bookHandler.UserClientHandler), bookHandler.CreateBook)
	group.Get("/list", middleware.APIUser(bookHandler.UserClientHandler), middleware.APIUserRolePermissions(bookHandler.UserClientHandler), bookHandler.GetBooks)
	group.Get("/:xid/detail", middleware.APIUser(bookHandler.UserClientHandler), middleware.APIUserRolePermissions(bookHandler.UserClientHandler), bookHandler.GetBook)
	group.Put("/:xid/update", middleware.APIUser(bookHandler.UserClientHandler), middleware.APIUserRolePermissions(bookHandler.UserClientHandler), bookHandler.UpdateBook)
	group.Delete("/:xid/delete", middleware.APIUser(bookHandler.UserClientHandler), middleware.APIUserRolePermissions(bookHandler.UserClientHandler), bookHandler.DeleteBook)

	group.Post("/:xid/borrow", middleware.APIUser(bookHandler.UserClientHandler), middleware.APIUserRolePermissions(bookHandler.UserClientHandler), bookHandler.BorrowBook)
	group.Post("/:xid/return", middleware.APIUser(bookHandler.UserClientHandler), middleware.APIUserRolePermissions(bookHandler.UserClientHandler), bookHandler.ReturnBook)
}
