// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package author

import (
	"context"

	pb_author "api-gateway/external/client/author-service/http/grpc/proto/author"
)

func (h *AuthorClientHandler) GetAuthor(ctx context.Context, in *pb_author.GetAuthorRequest) (*pb_author.GetAuthorResponse, error) {
	resp, err := h.Client.GetAuthor(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
