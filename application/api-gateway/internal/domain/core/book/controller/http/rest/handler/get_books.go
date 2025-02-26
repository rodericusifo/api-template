// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"api-gateway/internal/domain/core/book/controller/http/rest/request"
	"api-gateway/internal/pkg/types"
	"api-gateway/internal/pkg/util/response"
	"api-gateway/internal/pkg/util/serializer"
	"api-gateway/internal/pkg/util/validator"

	pb_book "api-gateway/external/client/book-service/http/grpc/proto/book"
)

func (h *BookHandler) GetBooks(ctx *fiber.Ctx) error {
	reqQuery := new(request.GetBooksRequestQuery)
	if err := validator.ValidateRequestQuery(ctx, reqQuery); err != nil {
		return err
	}

	getBooksRes, err := h.BookClientHandler.GetBooks(
		context.Background(),
		&pb_book.GetBooksRequest{
			Page:  reqQuery.Page,
			Limit: reqQuery.Limit,
		},
	)
	if err != nil {
		return err
	}

	meta := &types.Meta{
		CurrentPage:      getBooksRes.GetMeta().GetCurrentPage(),
		TotalDataPerPage: getBooksRes.GetMeta().GetTotalDataPerPage(),
		TotalPage:        getBooksRes.GetMeta().GetTotalPage(),
		TotalData:        getBooksRes.GetMeta().GetTotalData(),
	}
	getBooksResponse := serializer.SerializeBookProtosToBookResponses(getBooksRes.GetData())

	return ctx.Status(fiber.StatusOK).JSON(response.ResponseSuccess("get books success", getBooksResponse, meta))
}
