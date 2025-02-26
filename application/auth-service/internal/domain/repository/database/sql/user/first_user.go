// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package user

import (
	"auth-service/internal/domain/model/database/sql"
	"auth-service/internal/pkg/constant"
	"auth-service/internal/pkg/types"
	"auth-service/internal/pkg/util/builder"
)

func (r *UserDatabaseSQLRepository) FirstUser(query *types.QuerySQL) (*sql.User, error) {
	user := new(sql.User)

	q := r.Reader()

	if query != nil {
		q = builder.BuildQuerySQL(r.model.TableName(), q, query, constant.DialectDatabaseSQL(q.Dialector.Name()))
	}

	if err := q.Table(r.model.TableName()).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
