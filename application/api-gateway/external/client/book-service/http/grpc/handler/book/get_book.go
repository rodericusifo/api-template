// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package book

import (
	"context"

	pb_book "api-gateway/external/client/book-service/http/grpc/proto/book"
)

func (h *BookClientHandler) GetBook(ctx context.Context, in *pb_book.GetBookRequest) (*pb_book.GetBookResponse, error) {
	resp, err := h.Client.GetBook(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
