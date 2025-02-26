// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package book

import (
	"context"

	pb_book "api-gateway/external/client/book-service/http/grpc/proto/book"
)

func (h BookClientHandler) GetBooks(ctx context.Context, in *pb_book.GetBooksRequest) (*pb_book.GetBooksResponse, error) {
	resp, err := h.Client.GetBooks(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
