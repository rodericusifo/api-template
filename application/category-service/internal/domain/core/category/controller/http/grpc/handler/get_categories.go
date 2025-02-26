// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"context"

	"category-service/internal/domain/core/category/service/dto/input"
	"category-service/internal/pkg/util/serializer"

	pb_category "category-service/internal/proto/category"
)

func (h *CategoryHandler) GetCategories(ctx context.Context, req *pb_category.GetCategoriesRequest) (*pb_category.GetCategoriesResponse, error) {
	page, limit := int(req.GetPage()), int(req.GetLimit())

	categoryDtosRes, meta, err := h.CategoryService.GetCategories(&input.GetCategoriesDTO{
		Page:  &page,
		Limit: &limit,
	})
	if err != nil {
		return nil, err
	}
	return &pb_category.GetCategoriesResponse{
		Meta: &pb_category.Meta{
			CurrentPage:      int32(meta.CurrentPage),
			TotalDataPerPage: int32(meta.TotalDataPerPage),
			TotalData:        int32(meta.TotalData),
			TotalPage:        int32(meta.TotalPage),
		},
		Data: serializer.SerializeCategoryDTOsToCategoryProtos(categoryDtosRes),
	}, nil
}
