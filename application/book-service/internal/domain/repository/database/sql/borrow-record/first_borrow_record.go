// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package borrowrecord

import (
	"book-service/internal/domain/model/database/sql"
	"book-service/internal/pkg/constant"
	"book-service/internal/pkg/types"
	"book-service/internal/pkg/util/builder"
)

func (r *BorrowRecordDatabaseSQLRepository) FirstBorrowRecord(query *types.QuerySQL) (*sql.BorrowRecord, error) {
	borrowRecord := new(sql.BorrowRecord)

	q := r.Reader()

	if query != nil {
		q = builder.BuildQuerySQL(r.model.TableName(), q, query, constant.DialectDatabaseSQL(q.Dialector.Name()))
	}

	if err := q.Table(r.model.TableName()).First(borrowRecord).Error; err != nil {
		return nil, err
	}

	return borrowRecord, nil
}
