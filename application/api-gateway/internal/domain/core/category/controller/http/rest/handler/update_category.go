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

func (h *CategoryHandler) UpdateCategory(ctx *fiber.Ctx) error {
	reqBody := new(request.UpdateCategoryRequestBody)
	if err := validator.ValidateRequestBody(ctx, reqBody); err != nil {
		return err
	}

	reqParams := new(request.UpdateCategoryRequestParams)
	if err := validator.ValidateRequestParams(ctx, reqParams); err != nil {
		return err
	}

	if err := h.CategoryClientHandler.UpdateCategory(context.Background(), &pb_category.UpdateCategoryRequest{
		XID:         reqParams.XID,
		Name:        reqBody.Name,
		Description: reqBody.Description,
	}); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response.ResponseSuccess[any]("update category success", nil, nil))
}
