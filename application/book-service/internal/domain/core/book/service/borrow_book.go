// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package service

import (
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"book-service/internal/domain/core/book/service/dto/input"
	"book-service/internal/domain/model/database/sql"
	"book-service/internal/pkg/config"
	"book-service/internal/pkg/constant"
	"book-service/internal/pkg/types"
)

func (s *BookService) BorrowBook(payload *input.BorrowBookDTO) error {
	databaseSQL := config.GetDatabaseSQL(constant.POSTGRES)
	tx := databaseSQL.Writer.Begin()

	s.BookDatabaseSQLRepository.BeginTransaction(tx)
	defer s.BookDatabaseSQLRepository.EndTransaction()
	s.BorrowRecordDatabaseSQLRepository.BeginTransaction(tx)
	defer s.BorrowRecordDatabaseSQLRepository.EndTransaction()
	committed := false
	defer func() {
		if !committed {
			tx.Rollback()
		}
	}()

	firstBookModelRes, err := s.BookDatabaseSQLRepository.FirstBook(&types.QuerySQL{
		Searches: [][]types.SearchQuerySQLOperation{
			{
				{Field: "xid", Operator: "=", Value: payload.XID},
			},
		},
		Clauses: []types.ClauseExpression{
			clause.Locking{Strength: "UPDATE"},
		},
	})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return status.Error(codes.NotFound, "book not found")
		}
		return err
	}

	if firstBookModelRes.BorrowStock < 1 {
		return status.Error(codes.ResourceExhausted, "book not available")
	}

	borrowRecordSave := &sql.BorrowRecord{
		BorrowDate: time.Now(),
		BookID:     firstBookModelRes.ID,
		UserID:     payload.UserID,
	}

	err = s.BorrowRecordDatabaseSQLRepository.SaveBorrowRecord(borrowRecordSave)
	if err != nil {
		return err
	}

	bookSave := firstBookModelRes
	bookSave.BorrowStock -= 1

	err = s.BookDatabaseSQLRepository.SaveBook(bookSave)
	if err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}
	committed = true

	return nil
}
