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
)

func (h *AuthorHandler) UpdateAuthor(ctx *fiber.Ctx) error {
	reqBody := new(request.UpdateAuthorRequestBody)
	if err := validator.ValidateRequestBody(ctx, reqBody); err != nil {
		return err
	}

	reqParams := new(request.UpdateAuthorRequestParams)
	if err := validator.ValidateRequestParams(ctx, reqParams); err != nil {
		return err
	}

	if err := h.AuthorClientHandler.UpdateAuthor(context.Background(), &pb_author.UpdateAuthorRequest{
		XID:  reqParams.XID,
		Name: reqBody.Name,
		Bio:  reqBody.Bio,
	}); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response.ResponseSuccess[any]("update author success", nil, nil))
}
