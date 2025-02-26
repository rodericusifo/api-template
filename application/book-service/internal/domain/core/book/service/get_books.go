// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"book-service/internal/domain/core/book/service/dto/input"
	"book-service/internal/domain/core/book/service/dto/output"
	"book-service/internal/pkg/types"
	"book-service/internal/pkg/util/counter"
	"book-service/internal/pkg/util/definer"
	"book-service/internal/pkg/util/serializer"
)

func (s *BookService) GetBooks(payload *input.GetBooksDTO) (output.GetBooksDTO, *types.Meta, error) {
	page, limit := definer.DefinePaginationPageLimit(payload.Page, payload.Limit)

	bookListModelRes, err := s.BookDatabaseSQLRepository.FindBooks(&types.QuerySQL{
		Offset: counter.CountPaginationOffset(page, limit),
		Limit:  limit,
	})
	if err != nil {
		return nil, nil, err
	}
	countBookListModelRes := len(bookListModelRes)

	if len(bookListModelRes) < 1 {
		return nil, nil, status.Error(codes.NotFound, "books not found")
	}

	countBookAllModelRes, err := s.BookDatabaseSQLRepository.CountBooks(nil)
	if err != nil {
		return nil, nil, err
	}

	bookListDto := serializer.SerializeBooksToBookDTOs(bookListModelRes)

	meta := &types.Meta{
		CurrentPage:      int32(page),
		TotalDataPerPage: int32(countBookListModelRes),
		TotalData:        int32(countBookAllModelRes),
	}

	meta.TotalPage = int32(counter.CountPaginationTotalPage(int(meta.TotalDataPerPage), int(meta.TotalData)))

	return bookListDto, meta, nil
}
