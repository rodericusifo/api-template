// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package role

import (
	"auth-service/internal/domain/repository/database/sql/role"
)

type IRoleDatabaseSeederSQLRepository interface {
	SeedRoles()
}

type RoleDatabaseSeederSQLRepository struct {
	RoleDatabaseSQLRepository role.IRoleDatabaseSQLRepository
}

func InitRoleDatabaseSeederSQLRepository(roleDatabaseSQLRepository role.IRoleDatabaseSQLRepository) IRoleDatabaseSeederSQLRepository {
	return &RoleDatabaseSeederSQLRepository{
		RoleDatabaseSQLRepository: roleDatabaseSQLRepository,
	}
}
