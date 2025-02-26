// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"book-service/internal/domain/core/book/service/dto/input"

	pb_book "book-service/internal/proto/book"
)

func (h *BookHandler) CreateBook(ctx context.Context, req *pb_book.CreateBookRequest) (*emptypb.Empty, error) {
	publicationYear := req.GetPublicationYear()
	isbn := req.GetISBN()

	err := h.BookService.CreateBook(&input.CreateBookDTO{
		Title:           req.GetTitle(),
		ISBN:            &isbn,
		PublicationYear: &publicationYear,
		Stock:           req.GetStock(),
		AuthorID:        req.GetAuthorID(),
		CategoryID:      req.GetCategoryID(),
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
