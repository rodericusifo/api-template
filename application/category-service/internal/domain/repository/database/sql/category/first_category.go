// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package category

import (
	"category-service/internal/domain/model/database/sql"
	"category-service/internal/pkg/constant"
	"category-service/internal/pkg/types"
	"category-service/internal/pkg/util/builder"
)

func (r *CategoryDatabaseSQLRepository) FirstCategory(query *types.QuerySQL) (*sql.Category, error) {
	category := new(sql.Category)

	q := r.Reader()

	if query != nil {
		q = builder.BuildQuerySQL(r.model.TableName(), q, query, constant.DialectDatabaseSQL(q.Dialector.Name()))
	}

	if err := q.Table(r.model.TableName()).First(category).Error; err != nil {
		return nil, err
	}

	return category, nil
}
