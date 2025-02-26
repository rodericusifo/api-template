// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package borrowrecord

import (
	"book-service/internal/domain/model/database/sql"
)

func (r *BorrowRecordDatabaseSQLRepository) SaveBorrowRecord(payload *sql.BorrowRecord) error {
	borrowRecord := new(sql.BorrowRecord)

	q := r.Writer()

	if payload != nil {
		borrowRecord = payload
	}

	if err := q.Table(r.model.TableName()).Save(borrowRecord).Error; err != nil {
		return err
	}

	return nil
}
