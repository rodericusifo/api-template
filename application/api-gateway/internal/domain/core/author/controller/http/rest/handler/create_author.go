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

func (h *AuthorHandler) CreateAuthor(ctx *fiber.Ctx) error {
	reqBody := new(request.CreateAuthorRequestBody)
	if err := validator.ValidateRequestBody(ctx, reqBody); err != nil {
		return err
	}

	if err := h.AuthorClientHandler.CreateAuthor(context.Background(), &pb_author.CreateAuthorRequest{
		Name: reqBody.Name,
		Bio:  reqBody.Bio,
	}); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response.ResponseSuccess[any]("create author success", nil, nil))
}
