// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"api-gateway/internal/domain/core/category/controller/http/rest/request"
	"api-gateway/internal/pkg/types"
	"api-gateway/internal/pkg/util/response"
	"api-gateway/internal/pkg/util/serializer"
	"api-gateway/internal/pkg/util/validator"

	pb_category "api-gateway/external/client/category-service/http/grpc/proto/category"
)

func (h *CategoryHandler) GetCategories(ctx *fiber.Ctx) error {
	reqQuery := new(request.GetCategoriesRequestQuery)
	if err := validator.ValidateRequestQuery(ctx, reqQuery); err != nil {
		return err
	}

	getCategoriesRes, err := h.CategoryClientHandler.GetCategories(
		context.Background(),
		&pb_category.GetCategoriesRequest{
			Page:  reqQuery.Page,
			Limit: reqQuery.Limit,
		},
	)
	if err != nil {
		return err
	}

	meta := &types.Meta{
		CurrentPage:      getCategoriesRes.GetMeta().GetCurrentPage(),
		TotalDataPerPage: getCategoriesRes.GetMeta().GetTotalDataPerPage(),
		TotalPage:        getCategoriesRes.GetMeta().GetTotalPage(),
		TotalData:        getCategoriesRes.GetMeta().GetTotalData(),
	}
	getCategoriesResponse := serializer.SerializeCategoryProtosToCategoryResponses(getCategoriesRes.GetData())

	return ctx.Status(fiber.StatusOK).JSON(response.ResponseSuccess("get categories success", getCategoriesResponse, meta))
}
