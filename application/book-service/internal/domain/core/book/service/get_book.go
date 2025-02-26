// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"book-service/internal/domain/core/book/service/dto/input"
	"book-service/internal/domain/core/book/service/dto/output"
	"book-service/internal/pkg/types"
	"book-service/internal/pkg/util/serializer"
)

func (s *BookService) GetBook(payload *input.GetBookDTO) (output.GetBookDTO, error) {
	query := &types.QuerySQL{}
	if payload.XID != "" {
		query.Searches = [][]types.SearchQuerySQLOperation{
			{
				{Field: "xid", Operator: "=", Value: payload.XID},
			},
		}
	}
	if payload.AuthorID != nil {
		if *payload.AuthorID != 0 {
			query.Searches = [][]types.SearchQuerySQLOperation{
				{
					{Field: "author_id", Operator: "=", Value: *payload.AuthorID},
				},
			}
		}
	}
	if payload.CategoryID != nil {
		if *payload.CategoryID != 0 {
			query.Searches = [][]types.SearchQuerySQLOperation{
				{
					{Field: "category_id", Operator: "=", Value: *payload.CategoryID},
				},
			}
		}
	}

	bookModelRes, err := s.BookDatabaseSQLRepository.FirstBook(query)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(codes.NotFound, "book not found")
		}
		return nil, err
	}

	bookDto := serializer.SerializeBookToBookDTO(bookModelRes)

	return bookDto, nil
}
