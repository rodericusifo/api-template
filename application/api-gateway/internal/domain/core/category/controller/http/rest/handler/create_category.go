// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"api-gateway/internal/domain/core/category/controller/http/rest/request"
	"api-gateway/internal/pkg/util/response"
	"api-gateway/internal/pkg/util/validator"

	pb_category "api-gateway/external/client/category-service/http/grpc/proto/category"
)

func (h *CategoryHandler) CreateCategory(ctx *fiber.Ctx) error {
	reqBody := new(request.CreateCategoryRequestBody)
	if err := validator.ValidateRequestBody(ctx, reqBody); err != nil {
		return err
	}

	if err := h.CategoryClientHandler.CreateCategory(context.Background(), &pb_category.CreateCategoryRequest{
		Name:        reqBody.Name,
		Description: reqBody.Description,
	}); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response.ResponseSuccess[any]("create category success", nil, nil))
}
