// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package migration

import (
	"author-service/internal/domain/model/database/sql"
)

var (
	AutoMigrateModelList = []any{
		&sql.Author{},
	}
)
