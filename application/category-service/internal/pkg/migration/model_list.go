// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package migration

import (
	"category-service/internal/domain/model/database/sql"
)

var (
	AutoMigrateModelList = []any{
		&sql.Category{},
	}
)
