// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package migration

import (
	"book-service/internal/domain/model/database/sql"
)

var (
	AutoMigrateModelList = []any{
		&sql.Book{},
		&sql.BorrowRecord{},
	}
)
