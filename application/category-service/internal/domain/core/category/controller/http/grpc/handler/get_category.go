// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"context"

	"category-service/internal/domain/core/category/service/dto/input"
	"category-service/internal/pkg/util/serializer"

	pb_category "category-service/internal/proto/category"
)

func (h *CategoryHandler) GetCategory(ctx context.Context, req *pb_category.GetCategoryRequest) (*pb_category.GetCategoryResponse, error) {
	id := req.GetID()
	categoryDtoRes, err := h.CategoryService.GetCategory(&input.GetCategoryDTO{
		XID: req.GetXID(),
		ID:  &id,
	})
	if err != nil {
		return nil, err
	}
	return &pb_category.GetCategoryResponse{
		Data: serializer.SerializeCategoryDTOToCategoryProto(categoryDtoRes),
	}, nil
}
