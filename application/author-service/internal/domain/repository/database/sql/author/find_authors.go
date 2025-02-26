// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package author

import (
	"author-service/internal/domain/model/database/sql"
	"author-service/internal/pkg/constant"
	"author-service/internal/pkg/types"
	"author-service/internal/pkg/util/builder"
)

func (r *AuthorDatabaseSQLRepository) FindAuthors(query *types.QuerySQL) ([]*sql.Author, error) {
	authors := make([]*sql.Author, 0)

	q := r.Reader()

	if query != nil {
		q = builder.BuildQuerySQL(r.model.TableName(), q, query, constant.DialectDatabaseSQL(q.Dialector.Name()))
	}

	if err := q.Table(r.model.TableName()).Find(&authors).Error; err != nil {
		return nil, err
	}

	return authors, nil
}
