// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"api-gateway/internal/domain/core/author/controller/http/rest/request"
	"api-gateway/internal/pkg/types"
	"api-gateway/internal/pkg/util/response"
	"api-gateway/internal/pkg/util/serializer"
	"api-gateway/internal/pkg/util/validator"

	pb_author "api-gateway/external/client/author-service/http/grpc/proto/author"
)

func (h *AuthorHandler) GetAuthors(ctx *fiber.Ctx) error {
	reqQuery := new(request.GetAuthorsRequestQuery)
	if err := validator.ValidateRequestQuery(ctx, reqQuery); err != nil {
		return err
	}

	getAuthorsRes, err := h.AuthorClientHandler.GetAuthors(
		context.Background(),
		&pb_author.GetAuthorsRequest{
			Page:  reqQuery.Page,
			Limit: reqQuery.Limit,
		},
	)
	if err != nil {
		return err
	}

	meta := &types.Meta{
		CurrentPage:      getAuthorsRes.GetMeta().GetCurrentPage(),
		TotalDataPerPage: getAuthorsRes.GetMeta().GetTotalDataPerPage(),
		TotalPage:        getAuthorsRes.GetMeta().GetTotalPage(),
		TotalData:        getAuthorsRes.GetMeta().GetTotalData(),
	}
	getAuthorsResponse := serializer.SerializeAuthorProtosToAuthorResponses(getAuthorsRes.GetData())

	return ctx.Status(fiber.StatusOK).JSON(response.ResponseSuccess("get authors success", getAuthorsResponse, meta))
}
