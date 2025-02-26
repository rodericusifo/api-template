// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"context"

	"book-service/internal/domain/core/book/service/dto/input"
	"book-service/internal/pkg/util/serializer"

	pb_book "book-service/internal/proto/book"
)

func (h *BookHandler) GetBook(ctx context.Context, req *pb_book.GetBookRequest) (*pb_book.GetBookResponse, error) {
	authorID := req.GetAuthorID()
	categoryID := req.GetCategoryID()
	bookDtoRes, err := h.BookService.GetBook(&input.GetBookDTO{
		XID:        req.GetXID(),
		AuthorID:   &authorID,
		CategoryID: &categoryID,
	})
	if err != nil {
		return nil, err
	}
	return &pb_book.GetBookResponse{
		Data: serializer.SerializeBookDTOToBookProto(bookDtoRes),
	}, nil
}
