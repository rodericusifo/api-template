// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"book-service/internal/domain/core/book/service/dto/input"
	"book-service/internal/pkg/config"
	"book-service/internal/pkg/constant"
	"book-service/internal/pkg/types"
)

func (s *BookService) DeleteBook(payload *input.DeleteBookDTO) error {
	databaseSQL := config.GetDatabaseSQL(constant.POSTGRES)
	tx := databaseSQL.Writer.Begin()

	s.BookDatabaseSQLRepository.BeginTransaction(tx)
	defer s.BookDatabaseSQLRepository.EndTransaction()
	committed := false
	defer func() {
		if !committed {
			tx.Rollback()
		}
	}()

	bookModelRes, err := s.BookDatabaseSQLRepository.FirstBook(&types.QuerySQL{
		Selects: []types.SelectQuerySQLOperation{
			{Field: "id"},
		},
		Searches: [][]types.SearchQuerySQLOperation{
			{
				{Field: "xid", Operator: "=", Value: payload.XID},
			},
		},
	})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return status.Error(codes.NotFound, "book not found")
		}
		return err
	}

	bookModel := bookModelRes

	err = s.BookDatabaseSQLRepository.DeleteBook(bookModel)
	if err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}
	committed = true

	return nil
}
