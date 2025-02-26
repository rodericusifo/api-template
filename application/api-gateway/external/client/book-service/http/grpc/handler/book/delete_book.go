// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package book

import (
	"context"

	pb_book "api-gateway/external/client/book-service/http/grpc/proto/book"
)

func (h *BookClientHandler) DeleteBook(ctx context.Context, in *pb_book.DeleteBookRequest) error {
	_, err := h.Client.DeleteBook(ctx, in)
	if err != nil {
		return err
	}
	return nil
}
