// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package permission

import (
	"auth-service/internal/domain/repository/database/sql/permission"
)

type IPermissionDatabaseSeederSQLRepository interface {
	SeedPermissions()
}

type PermissionDatabaseSeederSQLRepository struct {
	PermissionDatabaseSQLRepository permission.IPermissionDatabaseSQLRepository
}

func InitPermissionDatabaseSeederSQLRepository(permissionDatabaseSQLRepository permission.IPermissionDatabaseSQLRepository) IPermissionDatabaseSeederSQLRepository {
	return &PermissionDatabaseSeederSQLRepository{
		PermissionDatabaseSQLRepository: permissionDatabaseSQLRepository,
	}
}
