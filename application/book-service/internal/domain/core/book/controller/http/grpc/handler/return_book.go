// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"book-service/internal/domain/core/book/service/dto/input"

	pb_book "book-service/internal/proto/book"
)

func (h *BookHandler) ReturnBook(ctx context.Context, req *pb_book.ReturnBookRequest) (*emptypb.Empty, error) {
	err := h.BookService.ReturnBook(&input.ReturnBookDTO{
		UserID: req.GetUserID(),
		XID:    req.GetXID(),
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
