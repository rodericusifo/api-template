// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package book

import (
	"gorm.io/gorm"

	"book-service/internal/domain/model/database/sql"
	"book-service/internal/pkg/config"
	"book-service/internal/pkg/interfaces"
	"book-service/internal/pkg/types"
)

type BookOperation struct {
	Query   *types.QuerySQL
	Payload *sql.Book
}

type BorrowRecordOperation struct {
	Query   *types.QuerySQL
	Payload *sql.BorrowRecord
}

type IBookDatabaseSQLRepository interface {
	interfaces.IDatabaseSQLRepository
	SaveBook(payload *sql.Book) error
	DeleteBook(payload *sql.Book) error
	FindBooks(query *types.QuerySQL) ([]*sql.Book, error)
	FirstBook(query *types.QuerySQL) (*sql.Book, error)
	CountBooks(query *types.QuerySQL) (int64, error)
}

type BookDatabaseSQLRepository struct {
	db    config.DatabaseSQL
	model sql.Book
	tx    *gorm.DB
}

func InitBookDatabaseSQLRepository(db config.DatabaseSQL) IBookDatabaseSQLRepository {
	return &BookDatabaseSQLRepository{
		db:    db,
		model: sql.Book{},
	}
}

func (r *BookDatabaseSQLRepository) BeginTransaction(tx *gorm.DB) {
	r.tx = tx
}

func (r *BookDatabaseSQLRepository) EndTransaction() {
	r.tx = nil
}

func (r *BookDatabaseSQLRepository) Writer() *gorm.DB {
	if r.tx != nil {
		return r.tx
	}
	return r.db.Writer
}

func (r *BookDatabaseSQLRepository) Reader() *gorm.DB {
	if r.tx != nil {
		return r.tx
	}
	return r.db.Reader
}
