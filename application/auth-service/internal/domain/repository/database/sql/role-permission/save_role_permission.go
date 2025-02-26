// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package rolepermission

import (
	"auth-service/internal/domain/model/database/sql"
)

func (r *RolePermissionDatabaseSQLRepository) SaveRolePermission(payload *sql.RolePermission) error {
	rolePermission := new(sql.RolePermission)

	q := r.Writer()

	if payload != nil {
		rolePermission = payload
	}

	if err := q.Table(r.model.TableName()).Save(rolePermission).Error; err != nil {
		return err
	}

	return nil
}
