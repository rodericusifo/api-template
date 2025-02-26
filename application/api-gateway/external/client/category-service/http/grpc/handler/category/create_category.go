// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package category

import (
	"context"

	pb_category "api-gateway/external/client/category-service/http/grpc/proto/category"
)

func (h *CategoryClientHandler) CreateCategory(ctx context.Context, in *pb_category.CreateCategoryRequest) error {
	_, err := h.Client.CreateCategory(ctx, in)
	if err != nil {
		return err
	}
	return nil
}
