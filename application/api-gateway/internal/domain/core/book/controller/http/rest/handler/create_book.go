// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"api-gateway/internal/domain/core/book/controller/http/rest/request"
	"api-gateway/internal/pkg/util/response"
	"api-gateway/internal/pkg/util/validator"

	pb_author "api-gateway/external/client/author-service/http/grpc/proto/author"
	pb_book "api-gateway/external/client/book-service/http/grpc/proto/book"
	pb_category "api-gateway/external/client/category-service/http/grpc/proto/category"
)

func (h *BookHandler) CreateBook(ctx *fiber.Ctx) error {
	reqBody := new(request.CreateBookRequestBody)
	if err := validator.ValidateRequestBody(ctx, reqBody); err != nil {
		return err
	}

	getAuthorRes, err := h.AuthorClientHandler.GetAuthor(context.Background(), &pb_author.GetAuthorRequest{
		XID: reqBody.AuthorXID,
	})
	if err != nil {
		return err
	}
	getCategoryRes, err := h.CategoryClientHandler.GetCategory(context.Background(), &pb_category.GetCategoryRequest{
		XID: reqBody.CategoryXID,
	})
	if err != nil {
		return err
	}

	if err := h.BookClientHandler.CreateBook(context.Background(), &pb_book.CreateBookRequest{
		Title:           reqBody.Title,
		ISBN:            reqBody.ISBN,
		PublicationYear: reqBody.PublicationYear,
		Stock:           reqBody.Stock,
		AuthorID:        getAuthorRes.GetData().GetID(),
		CategoryID:      getCategoryRes.GetData().GetID(),
	}); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response.ResponseSuccess[any]("create book success", nil, nil))
}
