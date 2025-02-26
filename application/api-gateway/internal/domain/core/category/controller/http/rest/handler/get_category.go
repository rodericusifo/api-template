// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"api-gateway/internal/domain/core/category/controller/http/rest/request"
	"api-gateway/internal/pkg/util/response"
	"api-gateway/internal/pkg/util/serializer"
	"api-gateway/internal/pkg/util/validator"

	pb_category "api-gateway/external/client/category-service/http/grpc/proto/category"
)

func (h *CategoryHandler) GetCategory(ctx *fiber.Ctx) error {
	reqParams := new(request.GetCategoryRequestParams)
	if err := validator.ValidateRequestParams(ctx, reqParams); err != nil {
		return err
	}

	getCategoryRes, err := h.CategoryClientHandler.GetCategory(context.Background(), &pb_category.GetCategoryRequest{
		XID: reqParams.XID,
	})
	if err != nil {
		return err
	}

	getCategoryResponse := serializer.SerializeCategoryProtoToCategoryResponse(getCategoryRes.GetData())

	return ctx.Status(fiber.StatusOK).JSON(response.ResponseSuccess("get category success", getCategoryResponse, nil))
}
