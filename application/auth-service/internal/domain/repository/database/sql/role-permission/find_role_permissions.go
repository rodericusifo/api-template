// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package rolepermission

import (
	"auth-service/internal/domain/model/database/sql"
	"auth-service/internal/pkg/constant"
	"auth-service/internal/pkg/types"
	"auth-service/internal/pkg/util/builder"
)

func (r *RolePermissionDatabaseSQLRepository) FindRolePermissions(query *types.QuerySQL) ([]*sql.RolePermission, error) {
	authors := make([]*sql.RolePermission, 0)

	q := r.Reader()

	if query != nil {
		q = builder.BuildQuerySQL(r.model.TableName(), q, query, constant.DialectDatabaseSQL(q.Dialector.Name()))
	}

	if err := q.Table(r.model.TableName()).Find(&authors).Error; err != nil {
		return nil, err
	}

	return authors, nil
}
