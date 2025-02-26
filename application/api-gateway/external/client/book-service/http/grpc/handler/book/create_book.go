// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package book

import (
	"context"

	pb_book "api-gateway/external/client/book-service/http/grpc/proto/book"
)

func (h *BookClientHandler) CreateBook(ctx context.Context, in *pb_book.CreateBookRequest) error {
	_, err := h.Client.CreateBook(ctx, in)
	if err != nil {
		return err
	}
	return nil
}
