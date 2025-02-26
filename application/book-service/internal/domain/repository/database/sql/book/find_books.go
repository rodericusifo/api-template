// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package book

import (
	"book-service/internal/domain/model/database/sql"
	"book-service/internal/pkg/constant"
	"book-service/internal/pkg/types"
	"book-service/internal/pkg/util/builder"
)

func (r *BookDatabaseSQLRepository) FindBooks(query *types.QuerySQL) ([]*sql.Book, error) {
	books := make([]*sql.Book, 0)

	q := r.Writer()

	if query != nil {
		q = builder.BuildQuerySQL(r.model.TableName(), q, query, constant.DialectDatabaseSQL(q.Dialector.Name()))
	}

	if err := q.Table(r.model.TableName()).Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}
