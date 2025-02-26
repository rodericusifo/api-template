// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package author

import (
	"context"

	pb_author "api-gateway/external/client/author-service/http/grpc/proto/author"
)

func (h *AuthorClientHandler) GetAuthors(ctx context.Context, in *pb_author.GetAuthorsRequest) (*pb_author.GetAuthorsResponse, error) {
	resp, err := h.Client.GetAuthors(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
