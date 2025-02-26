// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"category-service/internal/domain/core/category/service/dto/input"

	pb_category "category-service/internal/proto/category"
)

func (h *CategoryHandler) DeleteCategory(ctx context.Context, req *pb_category.DeleteCategoryRequest) (*emptypb.Empty, error) {
	err := h.CategoryService.DeleteCategory(&input.DeleteCategoryDTO{
		XID: req.GetXID(),
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
