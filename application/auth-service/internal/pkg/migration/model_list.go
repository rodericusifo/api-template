// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package migration

import (
	"auth-service/internal/domain/model/database/sql"
)

var (
	AutoMigrateModelList = []any{
		&sql.Permission{},
		&sql.Role{},
		&sql.RolePermission{},
		&sql.User{},
	}
)
