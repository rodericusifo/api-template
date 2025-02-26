// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package book

import (
	"context"

	pb_book "api-gateway/external/client/book-service/http/grpc/proto/book"
)

func (h *BookClientHandler) ReturnBook(ctx context.Context, in *pb_book.ReturnBookRequest) error {
	_, err := h.Client.ReturnBook(ctx, in)
	if err != nil {
		return err
	}
	return nil
}
