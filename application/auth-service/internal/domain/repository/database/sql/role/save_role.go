// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package role

import (
	"auth-service/internal/domain/model/database/sql"
)

func (r *RoleDatabaseSQLRepository) SaveRole(payload *sql.Role) error {
	role := new(sql.Role)

	q := r.Writer()

	if payload != nil {
		role = payload
	}

	if err := q.Table(r.model.TableName()).Save(role).Error; err != nil {
		return err
	}

	return nil
}
