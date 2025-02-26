// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package service

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"book-service/internal/domain/core/book/service/dto/input"
	"book-service/internal/pkg/config"
	"book-service/internal/pkg/constant"
	"book-service/internal/pkg/types"
)

func (s *BookService) UpdateBook(payload *input.UpdateBookDTO) error {
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

	bookModel := bookModelRes

	if payload.Title != nil {
		bookModelRes, err := s.BookDatabaseSQLRepository.FirstBook(&types.QuerySQL{
			Searches: [][]types.SearchQuerySQLOperation{
				{
					{Field: "title", Operator: "ILIKE", Value: payload.Title},
					{Field: "xid", Operator: "!=", Value: payload.XID},
				},
			},
		})
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}
		if bookModelRes != nil {
			return status.Error(codes.AlreadyExists, "book title already exist")
		}
		bookModel.Title = *payload.Title
	}
	if payload.ISBN != nil {
		bookModelRes, err := s.BookDatabaseSQLRepository.FirstBook(&types.QuerySQL{
			Searches: [][]types.SearchQuerySQLOperation{
				{
					{Field: "isbn", Operator: "=", Value: payload.ISBN},
					{Field: "xid", Operator: "!=", Value: payload.XID},
				},
			},
		})
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}
		if bookModelRes != nil {
			return status.Error(codes.AlreadyExists, "book isbn already exist")
		}
		bookModel.ISBN = payload.ISBN
	}

	bookModel.PublicationYear = payload.PublicationYear
	if payload.Stock != nil {
		stockAdjustment := *payload.Stock - bookModel.Stock
		borrowStockAdjustment := bookModel.BorrowStock + stockAdjustment
		if borrowStockAdjustment < 0 {
			minimumStock := bookModel.Stock - bookModel.BorrowStock
			return status.Error(codes.InvalidArgument, fmt.Sprintf("%d books borrowed, minimum stock is %d", minimumStock, minimumStock))
		}
		bookModel.Stock = *payload.Stock
		bookModel.BorrowStock = borrowStockAdjustment
	}
	if payload.AuthorID != nil {
		bookModel.AuthorID = *payload.AuthorID
	}
	if payload.CategoryID != nil {
		bookModel.CategoryID = *payload.CategoryID
	}

	err = s.BookDatabaseSQLRepository.SaveBook(bookModel)
	if err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}
	committed = true

	return nil
}
