// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"book-service/internal/domain/core/book/service/dto/input"

	pb_book "book-service/internal/proto/book"
)

func (h *BookHandler) DeleteBook(ctx context.Context, req *pb_book.DeleteBookRequest) (*emptypb.Empty, error) {
	err := h.BookService.DeleteBook(&input.DeleteBookDTO{
		XID: req.GetXID(),
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
