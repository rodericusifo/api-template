// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"api-gateway/internal/domain/core/author/controller/http/rest/request"
	"api-gateway/internal/pkg/util/response"
	"api-gateway/internal/pkg/util/validator"

	pb_author "api-gateway/external/client/author-service/http/grpc/proto/author"
	pb_book "api-gateway/external/client/book-service/http/grpc/proto/book"
)

func (h *AuthorHandler) DeleteAuthor(ctx *fiber.Ctx) error {
	reqParams := new(request.DeleteAuthorRequestParams)
	if err := validator.ValidateRequestParams(ctx, reqParams); err != nil {
		return err
	}

	getAuthorRes, err := h.AuthorClientHandler.GetAuthor(context.Background(), &pb_author.GetAuthorRequest{
		XID: reqParams.XID,
	})
	if err != nil {
		return err
	}

	authorID := getAuthorRes.GetData().GetID()
	getBookRes, err := h.BookClientHandler.GetBook(context.Background(), &pb_book.GetBookRequest{
		AuthorID: &authorID,
	})
	if err != nil && err.Error() != "book not found" {
		return err
	}
	if getBookRes.GetData().GetID() != 0 {
		return fiber.NewError(fiber.StatusConflict, "Cannot delete author because it is referenced by a book")
	}

	if err := h.AuthorClientHandler.DeleteAuthor(context.Background(), &pb_author.DeleteAuthorRequest{
		XID: reqParams.XID,
	}); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response.ResponseSuccess[any]("delete author success", nil, nil))
}
