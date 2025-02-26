// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"api-gateway/internal/domain/core/author/controller/http/rest/request"
	"api-gateway/internal/pkg/util/response"
	"api-gateway/internal/pkg/util/serializer"
	"api-gateway/internal/pkg/util/validator"

	pb_author "api-gateway/external/client/author-service/http/grpc/proto/author"
)

func (h *AuthorHandler) GetAuthor(ctx *fiber.Ctx) error {
	reqParams := new(request.GetAuthorRequestParams)
	if err := validator.ValidateRequestParams(ctx, reqParams); err != nil {
		return err
	}

	getAuthorRes, err := h.AuthorClientHandler.GetAuthor(context.Background(), &pb_author.GetAuthorRequest{
		XID: reqParams.XID,
	})
	if err != nil {
		return err
	}

	getAuthorResponse := serializer.SerializeAuthorProtoToAuthorResponse(getAuthorRes.GetData())

	return ctx.Status(fiber.StatusOK).JSON(response.ResponseSuccess("get author success", getAuthorResponse, nil))
}
