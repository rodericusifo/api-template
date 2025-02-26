// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package permission

import (
	"auth-service/internal/domain/model/database/sql"
)

func (r *PermissionDatabaseSQLRepository) SavePermission(payload *sql.Permission) error {
	permission := new(sql.Permission)

	q := r.Writer()

	if payload != nil {
		permission = payload
	}

	if err := q.Table(r.model.TableName()).Save(permission).Error; err != nil {
		return err
	}

	return nil
}
