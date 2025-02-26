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

func (h *BookHandler) UpdateBook(ctx *fiber.Ctx) error {
	reqBody := new(request.UpdateBookRequestBody)
	if err := validator.ValidateRequestBody(ctx, reqBody); err != nil {
		return err
	}

	reqParams := new(request.UpdateBookRequestParams)
	if err := validator.ValidateRequestParams(ctx, reqParams); err != nil {
		return err
	}

	grpcReq := &pb_book.UpdateBookRequest{
		XID:             reqParams.XID,
		Title:           reqBody.Title,
		ISBN:            reqBody.ISBN,
		PublicationYear: reqBody.PublicationYear,
		Stock:           reqBody.Stock,
	}

	if reqBody.AuthorXID != nil {
		getAuthorRes, err := h.AuthorClientHandler.GetAuthor(context.Background(), &pb_author.GetAuthorRequest{
			XID: *reqBody.AuthorXID,
		})
		if err != nil {
			return err
		}
		authorID := getAuthorRes.GetData().GetID()
		grpcReq.AuthorID = &authorID
	}
	if reqBody.CategoryXID != nil {
		getCategoryRes, err := h.CategoryClientHandler.GetCategory(context.Background(), &pb_category.GetCategoryRequest{
			XID: *reqBody.CategoryXID,
		})
		if err != nil {
			return err
		}
		categoryID := getCategoryRes.GetData().GetID()
		grpcReq.CategoryID = &categoryID
	}

	if err := h.BookClientHandler.UpdateBook(context.Background(), grpcReq); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response.ResponseSuccess[any]("update book success", nil, nil))
}
