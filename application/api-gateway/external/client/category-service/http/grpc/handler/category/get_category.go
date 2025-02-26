// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package category

import (
	"context"

	pb_category "api-gateway/external/client/category-service/http/grpc/proto/category"
)

func (h *CategoryClientHandler) GetCategory(ctx context.Context, in *pb_category.GetCategoryRequest) (*pb_category.GetCategoryResponse, error) {
	resp, err := h.Client.GetCategory(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
