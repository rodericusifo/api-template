// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"category-service/internal/domain/core/category/service/dto/input"

	pb_category "category-service/internal/proto/category"
)

func (h *CategoryHandler) CreateCategory(ctx context.Context, req *pb_category.CreateCategoryRequest) (*emptypb.Empty, error) {
	description := req.GetDescription()
	err := h.CategoryService.CreateCategory(&input.CreateCategoryDTO{
		Name:        req.GetName(),
		Description: &description,
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
