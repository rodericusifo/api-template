// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"book-service/internal/domain/core/book/service/dto/input"

	pb_book "book-service/internal/proto/book"
)

func (h *BookHandler) UpdateBook(ctx context.Context, req *pb_book.UpdateBookRequest) (*emptypb.Empty, error) {
	publicationYear := req.GetPublicationYear()
	stock := req.GetStock()

	authorID := req.GetAuthorID()
	categoryID := req.GetCategoryID()

	title := req.GetTitle()
	isbn := req.GetISBN()

	dto := &input.UpdateBookDTO{
		XID:             req.GetXID(),
		Title:           &title,
		ISBN:            &isbn,
		PublicationYear: &publicationYear,
		Stock:           &stock,
		AuthorID:        &authorID,
		CategoryID:      &categoryID,
	}

	err := h.BookService.UpdateBook(dto)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
