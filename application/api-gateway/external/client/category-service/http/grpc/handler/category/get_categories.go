// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package category

import (
	"context"

	pb_category "api-gateway/external/client/category-service/http/grpc/proto/category"
)

func (h *CategoryClientHandler) GetCategories(ctx context.Context, in *pb_category.GetCategoriesRequest) (*pb_category.GetCategoriesResponse, error) {
	resp, err := h.Client.GetCategories(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
