// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package author

import (
	"context"

	pb_author "api-gateway/external/client/author-service/http/grpc/proto/author"
)

func (h *AuthorClientHandler) UpdateAuthor(ctx context.Context, in *pb_author.UpdateAuthorRequest) error {
	_, err := h.Client.UpdateAuthor(ctx, in)
	if err != nil {
		return err
	}
	return nil
}
