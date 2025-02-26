// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package author

import (
	"author-service/internal/pkg/constant"
	"author-service/internal/pkg/types"
	"author-service/internal/pkg/util/builder"
)

func (r *AuthorDatabaseSQLRepository) CountAuthors(query *types.QuerySQL) (int64, error) {
	count := int64(0)

	q := r.Reader()

	defaultQuery := &types.QuerySQL{
		Selects: []types.SelectQuerySQLOperation{
			{Field: "id"},
		},
	}

	if query != nil {
		query.Selects = append(query.Selects, defaultQuery.Selects...)
		q = builder.BuildQuerySQL(r.model.TableName(), q, query, constant.DialectDatabaseSQL(q.Dialector.Name()))
	} else {
		q = builder.BuildQuerySQL(r.model.TableName(), q, defaultQuery, constant.DialectDatabaseSQL(q.Dialector.Name()))
	}

	if err := q.Table(r.model.TableName()).Count(&count).Error; err != nil {
		return count, err
	}

	return count, nil
}
