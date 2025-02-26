// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"github.com/gofiber/fiber/v2"

	"api-gateway/internal/pkg/middleware"

	external_client_authservice_http_grpc_handler_user "api-gateway/external/client/auth-service/http/grpc/handler/user"
	external_client_bookservice_http_grpc_handler_book "api-gateway/external/client/book-service/http/grpc/handler/book"
	external_client_categoryservice_http_grpc_handler_category "api-gateway/external/client/category-service/http/grpc/handler/category"
)

type CategoryHandler struct {
	UserClientHandler     external_client_authservice_http_grpc_handler_user.IUserClientHandler
	CategoryClientHandler external_client_categoryservice_http_grpc_handler_category.ICategoryClientHandler
	BookClientHandler     external_client_bookservice_http_grpc_handler_book.IBookClientHandler
}

func InitCategoryHandler(
	userClientHandler external_client_authservice_http_grpc_handler_user.IUserClientHandler,
	categoryClientHandler external_client_categoryservice_http_grpc_handler_category.ICategoryClientHandler,
	bookClientHandler external_client_bookservice_http_grpc_handler_book.IBookClientHandler,
) *CategoryHandler {
	return &CategoryHandler{
		UserClientHandler:     userClientHandler,
		CategoryClientHandler: categoryClientHandler,
		BookClientHandler:     bookClientHandler,
	}
}

func (categoryHandler *CategoryHandler) Mount(group fiber.Router) {
	group.Post("/create", middleware.APIUser(categoryHandler.UserClientHandler), middleware.APIUserRolePermissions(categoryHandler.UserClientHandler), categoryHandler.CreateCategory)
	group.Get("/list", middleware.APIUser(categoryHandler.UserClientHandler), middleware.APIUserRolePermissions(categoryHandler.UserClientHandler), categoryHandler.GetCategories)
	group.Get("/:xid/detail", middleware.APIUser(categoryHandler.UserClientHandler), middleware.APIUserRolePermissions(categoryHandler.UserClientHandler), categoryHandler.GetCategory)
	group.Put("/:xid/update", middleware.APIUser(categoryHandler.UserClientHandler), middleware.APIUserRolePermissions(categoryHandler.UserClientHandler), categoryHandler.UpdateCategory)
	group.Delete("/:xid/delete", middleware.APIUser(categoryHandler.UserClientHandler), middleware.APIUserRolePermissions(categoryHandler.UserClientHandler), categoryHandler.DeleteCategory)
}
