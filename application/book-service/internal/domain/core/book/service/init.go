// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package service

import (
	"book-service/internal/domain/core/book/service/dto/input"
	"book-service/internal/domain/core/book/service/dto/output"
	"book-service/internal/domain/repository/database/sql/book"
	"book-service/internal/pkg/types"

	borrowrecord "book-service/internal/domain/repository/database/sql/borrow-record"
)

type IBookService interface {
	CreateBook(payload *input.CreateBookDTO) error
	UpdateBook(payload *input.UpdateBookDTO) error
	DeleteBook(payload *input.DeleteBookDTO) error
	GetBooks(payload *input.GetBooksDTO) (output.GetBooksDTO, *types.Meta, error)
	GetBook(payload *input.GetBookDTO) (output.GetBookDTO, error)

	BorrowBook(payload *input.BorrowBookDTO) error
	ReturnBook(payload *input.ReturnBookDTO) error
}

type BookService struct {
	BookDatabaseSQLRepository         book.IBookDatabaseSQLRepository
	BorrowRecordDatabaseSQLRepository borrowrecord.IBorrowRecordDatabaseSQLRepository
}

func InitBookService(
	bookDatabaseSQLRepository book.IBookDatabaseSQLRepository,
	borrowRecordDatabaseSQLRepository borrowrecord.IBorrowRecordDatabaseSQLRepository,
) IBookService {
	return &BookService{
		BookDatabaseSQLRepository:         bookDatabaseSQLRepository,
		BorrowRecordDatabaseSQLRepository: borrowRecordDatabaseSQLRepository,
	}
}
