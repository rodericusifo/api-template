// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"api-gateway/internal/domain/core/category/controller/http/rest/request"
	"api-gateway/internal/pkg/util/response"
	"api-gateway/internal/pkg/util/validator"

	pb_book "api-gateway/external/client/book-service/http/grpc/proto/book"
	pb_category "api-gateway/external/client/category-service/http/grpc/proto/category"
)

func (h *CategoryHandler) DeleteCategory(ctx *fiber.Ctx) error {
	reqParams := new(request.DeleteCategoryRequestParams)
	if err := validator.ValidateRequestParams(ctx, reqParams); err != nil {
		return err
	}

	getCategoryRes, err := h.CategoryClientHandler.GetCategory(context.Background(), &pb_category.GetCategoryRequest{
		XID: reqParams.XID,
	})
	if err != nil {
		return err
	}

	categoryID := getCategoryRes.GetData().GetID()
	getBookRes, err := h.BookClientHandler.GetBook(context.Background(), &pb_book.GetBookRequest{
		CategoryID: &categoryID,
	})
	if err != nil && err.Error() != "book not found" {
		return err
	}
	if getBookRes.GetData().GetID() != 0 {
		return fiber.NewError(fiber.StatusConflict, "Cannot delete category because it is referenced by a book")
	}

	if err := h.CategoryClientHandler.DeleteCategory(context.Background(), &pb_category.DeleteCategoryRequest{
		XID: reqParams.XID,
	}); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response.ResponseSuccess[any]("delete category success", nil, nil))
}
