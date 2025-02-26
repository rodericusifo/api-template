// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package book

import (
	"book-service/internal/domain/model/database/sql"
)

func (r *BookDatabaseSQLRepository) DeleteBook(payload *sql.Book) error {
	book := new(sql.Book)

	q := r.Writer()

	if payload != nil {
		book = payload
	}

	if err := q.Table(r.model.TableName()).Unscoped().Delete(book).Error; err != nil {
		return err
	}

	return nil
}
