// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"api-gateway/internal/domain/core/book/controller/http/rest/request"
	"api-gateway/internal/pkg/util/response"
	"api-gateway/internal/pkg/util/serializer"
	"api-gateway/internal/pkg/util/validator"

	pb_author "api-gateway/external/client/author-service/http/grpc/proto/author"
	pb_book "api-gateway/external/client/book-service/http/grpc/proto/book"
	pb_category "api-gateway/external/client/category-service/http/grpc/proto/category"
)

func (h *BookHandler) GetBook(ctx *fiber.Ctx) error {
	reqParams := new(request.GetBookRequestParams)
	if err := validator.ValidateRequestParams(ctx, reqParams); err != nil {
		return err
	}

	getBookRes, err := h.BookClientHandler.GetBook(context.Background(), &pb_book.GetBookRequest{
		XID: reqParams.XID,
	})
	if err != nil {
		return err
	}

	authorID := getBookRes.GetData().GetAuthorID()
	getAuthorRes, _ := h.AuthorClientHandler.GetAuthor(context.Background(), &pb_author.GetAuthorRequest{
		ID: &authorID,
	})

	categoryID := getBookRes.GetData().GetCategoryID()
	getCategoryRes, _ := h.CategoryClientHandler.GetCategory(context.Background(), &pb_category.GetCategoryRequest{
		ID: &categoryID,
	})

	getBookResponse := serializer.SerializeBookProtoToBookResponse(getBookRes.GetData(), getAuthorRes.GetData(), getCategoryRes.GetData())

	return ctx.Status(fiber.StatusOK).JSON(response.ResponseSuccess("get book success", getBookResponse, nil))
}
