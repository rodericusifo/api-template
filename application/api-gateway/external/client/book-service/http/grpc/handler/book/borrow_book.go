// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package book

import (
	"context"

	pb_book "api-gateway/external/client/book-service/http/grpc/proto/book"
)

func (h *BookClientHandler) BorrowBook(ctx context.Context, in *pb_book.BorrowBookRequest) error {
	_, err := h.Client.BorrowBook(ctx, in)
	if err != nil {
		return err
	}
	return nil
}
