// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package author

import (
	"author-service/internal/domain/model/database/sql"
)

func (r *AuthorDatabaseSQLRepository) DeleteAuthor(payload *sql.Author) error {
	author := new(sql.Author)

	q := r.Writer()

	if payload != nil {
		author = payload
	}

	if err := q.Table(r.model.TableName()).Unscoped().Delete(author).Error; err != nil {
		return err
	}

	return nil
}
