// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"context"

	"book-service/internal/domain/core/book/service/dto/input"
	"book-service/internal/pkg/util/serializer"

	pb_book "book-service/internal/proto/book"
)

func (h *BookHandler) GetBooks(ctx context.Context, req *pb_book.GetBooksRequest) (*pb_book.GetBooksResponse, error) {
	page, limit := int(req.GetPage()), int(req.GetLimit())

	bookDtosRes, meta, err := h.BookService.GetBooks(&input.GetBooksDTO{
		Page:  &page,
		Limit: &limit,
	})
	if err != nil {
		return nil, err
	}
	return &pb_book.GetBooksResponse{
		Meta: &pb_book.Meta{
			CurrentPage:      int32(meta.CurrentPage),
			TotalDataPerPage: int32(meta.TotalDataPerPage),
			TotalData:        int32(meta.TotalData),
			TotalPage:        int32(meta.TotalPage),
		},
		Data: serializer.SerializeBookDTOsToBookProtos(bookDtosRes),
	}, nil
}
