// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"category-service/internal/domain/core/category/service/dto/input"

	pb_category "category-service/internal/proto/category"
)

func (h *CategoryHandler) UpdateCategory(ctx context.Context, req *pb_category.UpdateCategoryRequest) (*emptypb.Empty, error) {
	name := req.GetName()
	description := req.GetDescription()
	err := h.CategoryService.UpdateCategory(&input.UpdateCategoryDTO{
		XID:         req.GetXID(),
		Name:        &name,
		Description: &description,
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
