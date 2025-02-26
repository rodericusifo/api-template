// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"api-gateway/internal/domain/core/book/controller/http/rest/request"
	"api-gateway/internal/pkg/util/getter"
	"api-gateway/internal/pkg/util/response"
	"api-gateway/internal/pkg/util/validator"

	pb_book "api-gateway/external/client/book-service/http/grpc/proto/book"
)

func (h *BookHandler) ReturnBook(ctx *fiber.Ctx) error {
	reqUser := getter.GetRequestUser(ctx)

	reqParams := new(request.ReturnBookRequestParams)
	if err := validator.ValidateRequestParams(ctx, reqParams); err != nil {
		return err
	}

	if err := h.BookClientHandler.ReturnBook(context.Background(), &pb_book.ReturnBookRequest{
		UserID: reqUser.ID,
		XID:    reqParams.XID,
	}); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response.ResponseSuccess[any]("return book success", nil, nil))
}
