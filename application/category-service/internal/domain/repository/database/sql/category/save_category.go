// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package category

import (
	"category-service/internal/domain/model/database/sql"
)

func (r *CategoryDatabaseSQLRepository) SaveCategory(payload *sql.Category) error {
	category := new(sql.Category)

	q := r.Writer()

	if payload != nil {
		category = payload
	}

	if err := q.Table(r.model.TableName()).Save(category).Error; err != nil {
		return err
	}

	return nil
}
