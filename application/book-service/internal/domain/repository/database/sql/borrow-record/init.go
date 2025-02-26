// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package borrowrecord

import (
	"gorm.io/gorm"

	"book-service/internal/domain/model/database/sql"
	"book-service/internal/pkg/config"
	"book-service/internal/pkg/interfaces"
	"book-service/internal/pkg/types"
)

type IBorrowRecordDatabaseSQLRepository interface {
	interfaces.IDatabaseSQLRepository
	SaveBorrowRecord(payload *sql.BorrowRecord) error
	FirstBorrowRecord(query *types.QuerySQL) (*sql.BorrowRecord, error)
}

type BorrowRecordDatabaseSQLRepository struct {
	db    config.DatabaseSQL
	model sql.BorrowRecord
	tx    *gorm.DB
}

func InitBorrowRecordDatabaseSQLRepository(db config.DatabaseSQL) IBorrowRecordDatabaseSQLRepository {
	return &BorrowRecordDatabaseSQLRepository{
		db:    db,
		model: sql.BorrowRecord{},
	}
}

func (r *BorrowRecordDatabaseSQLRepository) BeginTransaction(tx *gorm.DB) {
	r.tx = tx
}

func (r *BorrowRecordDatabaseSQLRepository) EndTransaction() {
	r.tx = nil
}

func (r *BorrowRecordDatabaseSQLRepository) Writer() *gorm.DB {
	if r.tx != nil {
		return r.tx
	}
	return r.db.Writer
}

func (r *BorrowRecordDatabaseSQLRepository) Reader() *gorm.DB {
	if r.tx != nil {
		return r.tx
	}
	return r.db.Reader
}
