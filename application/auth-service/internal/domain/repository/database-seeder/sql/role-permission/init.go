// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package rolepermission

import (
	"auth-service/internal/domain/repository/database/sql/permission"
	"auth-service/internal/domain/repository/database/sql/role"
	rolepermission "auth-service/internal/domain/repository/database/sql/role-permission"
)

type IRolePermissionDatabaseSeederSQLRepository interface {
	SeedRolePermissions()
}

type RolePermissionDatabaseSeederSQLRepository struct {
	RoleDatabaseSQLRepository           role.IRoleDatabaseSQLRepository
	PermissionDatabaseSQLRepository     permission.IPermissionDatabaseSQLRepository
	RolePermissionDatabaseSQLRepository rolepermission.IRolePermissionDatabaseSQLRepository
}

func InitRolePermissionDatabaseSeederSQLRepository(
	roleDatabaseSQLRepository role.IRoleDatabaseSQLRepository,
	permissionDatabaseSQLRepository permission.IPermissionDatabaseSQLRepository,
	rolePermissionDatabaseSQLRepository rolepermission.IRolePermissionDatabaseSQLRepository,
) IRolePermissionDatabaseSeederSQLRepository {
	return &RolePermissionDatabaseSeederSQLRepository{
		RoleDatabaseSQLRepository:           roleDatabaseSQLRepository,
		PermissionDatabaseSQLRepository:     permissionDatabaseSQLRepository,
		RolePermissionDatabaseSQLRepository: rolePermissionDatabaseSQLRepository,
	}
}
