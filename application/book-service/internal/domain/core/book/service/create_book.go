// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"book-service/internal/domain/core/book/service/dto/input"
	"book-service/internal/domain/model/database/sql"
	"book-service/internal/pkg/types"
)

func (s *BookService) CreateBook(payload *input.CreateBookDTO) error {
	query := &types.QuerySQL{
		Selects: []types.SelectQuerySQLOperation{
			{Field: "id"},
		},
		Searches: [][]types.SearchQuerySQLOperation{
			{
				{Field: "title", Operator: "ILIKE", Value: payload.Title},
			},
		},
	}
	if payload.ISBN != nil {
		query.Searches = append(query.Searches, []types.SearchQuerySQLOperation{
			{Field: "isbn", Operator: "=", Value: payload.ISBN},
		})
	}

	bookModelRes, err := s.BookDatabaseSQLRepository.FirstBook(query)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if bookModelRes != nil {
		return status.Error(codes.AlreadyExists, "book already exist")
	}

	bookModel := &sql.Book{
		Title:           payload.Title,
		ISBN:            payload.ISBN,
		PublicationYear: payload.PublicationYear,
		Stock:           payload.Stock,
		BorrowStock:     payload.Stock,
		AuthorID:        payload.AuthorID,
		CategoryID:      payload.CategoryID,
	}
	err = s.BookDatabaseSQLRepository.SaveBook(bookModel)
	if err != nil {
		return err
	}

	return nil
}
