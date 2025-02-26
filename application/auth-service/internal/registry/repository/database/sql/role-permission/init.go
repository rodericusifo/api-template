// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package rolepermission

import (
	"auth-service/internal/pkg/config"
	"auth-service/internal/pkg/constant"

	internal_domain_repository_database_sql_rolepermission "auth-service/internal/domain/repository/database/sql/role-permission"
)

func RolePermissionDatabaseSQLRepository(dialect constant.DialectDatabaseSQL) internal_domain_repository_database_sql_rolepermission.IRolePermissionDatabaseSQLRepository {
	databaseSQL := config.GetDatabaseSQL(dialect)
	iRolePermissionDatabaseSQLRepository := internal_domain_repository_database_sql_rolepermission.InitRolePermissionDatabaseSQLRepository(databaseSQL)
	return iRolePermissionDatabaseSQLRepository
}
